package utils

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

type FileType int

const (
	FILE FileType = iota
	DIR
)

type FileInfo struct {
	FType FileType `json:"ftype"`
	Size  int      `json:"size"`
	Name  string   `json:"name"`
}

func parseLSOutput(output string) ([]FileInfo, error) {
	var files []FileInfo
	scanner := bufio.NewScanner(bytes.NewBufferString(output))

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) >= 9 {
			fInfo := FileInfo{}
			flag := []byte(fields[0])[0]
			if flag == 'd' {
				fInfo.FType = DIR
			} else if flag == '-' {
				fInfo.FType = FILE
				size, err := strconv.Atoi(fields[4])
				if err != nil {
					return files, err
				}

				fInfo.Size = size
			} else {
				continue
			}

			t := fields[7]
			index := strings.Index(line, t) + len(t) + 1
			if index > 0 {
				fInfo.Name = line[index:]
			}

			if fInfo.Name == "." || fInfo.Name == ".." {
				continue
			}

			files = append(files, fInfo)
		}
	}

	return files, nil
}

type DockerTermSize struct {
	Width  uint `json:"width"`
	Height uint `json:"height"`
}

type LogsFun func(reader io.ReadCloser) error
type ExecAttachFun func(conn net.Conn, sizeChan chan DockerTermSize) error

type DockerApi struct {
	uri string
}

func (d *DockerApi) SetURI(uri string) {
	d.uri = uri
}

func (d *DockerApi) SetTcp(host string, port int) {
	d.uri = fmt.Sprintf("tcp://%s:%d", host, port)
}

func (d *DockerApi) SetUnix(socketFile string) {
	d.uri = fmt.Sprintf("unix://%s", socketFile)
}

func (d *DockerApi) Conn() (*client.Client, error) {
	return client.NewClientWithOpts(client.WithHost(d.uri), client.WithAPIVersionNegotiation())
}

func (d *DockerApi) Version() (*types.Version, error) {
	cli, err := d.Conn()
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	ver, err := cli.ServerVersion(context.Background())
	if err != nil {
		return nil, err
	}

	return &ver, err
}

type ContainerInfo struct {
	Summary *container.Summary
	Inspect *container.InspectResponse
}

func (d *DockerApi) ContainerList() ([]*ContainerInfo, error) {
	cli, err := d.Conn()
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	ctx := context.Background()
	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	cs := []*ContainerInfo{}

	for _, c := range containers {
		info := ContainerInfo{}
		info.Summary = &c

		inspect, err := cli.ContainerInspect(ctx, c.ID)
		if err != nil {
			return nil, err
		}

		info.Inspect = &inspect
		cs = append(cs, &info)
	}

	return cs, err
}

func (d *DockerApi) ContainerDir(id string, dir string) ([]FileInfo, error) {
	ret, err := d.ExecCmd(id, []string{"ls", "-al", dir})
	if err != nil {
		return nil, err
	}

	return parseLSOutput(ret)
}

func (d *DockerApi) StartContainer(id string) error {
	cli, err := d.Conn()
	if err != nil {
		return err
	}
	defer cli.Close()
	return cli.ContainerStart(context.Background(), id, container.StartOptions{})
}

func (d *DockerApi) StopContainer(id string) error {
	cli, err := d.Conn()
	if err != nil {
		return err
	}
	defer cli.Close()

	return cli.ContainerStop(context.Background(), id, container.StopOptions{})
}

func (d *DockerApi) RestartContainer(id string) error {
	cli, err := d.Conn()
	if err != nil {
		return err
	}
	defer cli.Close()

	return cli.ContainerRestart(context.Background(), id, container.StopOptions{})
}

func (d *DockerApi) DeleteContainer(id string) error {
	cli, err := d.Conn()
	if err != nil {
		return err
	}
	defer cli.Close()

	return cli.ContainerRemove(context.Background(), id,
		container.RemoveOptions{Force: true})
}

func (d *DockerApi) execCmd(ctx context.Context, cli *client.Client, id string, cmd []string) (string, error) {
	cfg := container.ExecOptions{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          false,
		Cmd:          cmd,
		Privileged:   true,
	}

	execID, err := cli.ContainerExecCreate(ctx, id, cfg)
	if err != nil {
		return "", err
	}

	resp, _ := cli.ContainerExecAttach(ctx, execID.ID, container.ExecAttachOptions{
		Tty: false,
	})
	defer resp.Close()

	buf := bytes.NewBufferString("")
	io.Copy(buf, resp.Reader)

	inspect, err := cli.ContainerExecInspect(ctx, execID.ID)
	if err != nil {
		return buf.String(), err
	}

	if inspect.ExitCode != 0 {
		return buf.String(), fmt.Errorf("cmd error：%d", inspect.ExitCode)
	}

	return buf.String(), err
}

func (d *DockerApi) ExecCmd(id string, cmd []string) (string, error) {
	cli, err := d.Conn()
	if err != nil {
		return "", err
	}
	defer cli.Close()

	ctx := context.Background()

	return d.execCmd(ctx, cli, id, cmd)
}

func (d *DockerApi) DockerTerm(id string) string {
	cmd := "/bin/bash"

	_, err := d.ExecCmd(id, []string{"ls", cmd})
	if err != nil {
		cmd = "sh"
	} else {
		cmd = "bash"
	}

	return cmd
}

func (d *DockerApi) DockerInspect(id string) (*container.InspectResponse, error) {
	cli, err := d.Conn()
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	ctx := context.Background()

	info, err := cli.ContainerInspect(ctx, id)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (d *DockerApi) ExecAttach(id string, workDir string, fun ExecAttachFun) error {
	cli, err := d.Conn()
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx := context.Background()

	env := []string{}
	env = append(env, "TERM=xterm")
	env = append(env, "LANG=zh_CN.UTF-8")

	cfg := container.ExecOptions{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
		Cmd:          []string{d.DockerTerm(id)},
		Privileged:   true,
		Env:          env,
	}

	if len(workDir) != 0 {
		cfg.WorkingDir = workDir
	}

	execID, err := cli.ContainerExecCreate(ctx, id, cfg)
	if err != nil {
		return err
	}

	resp, err := cli.ContainerExecAttach(ctx, execID.ID, container.ExecAttachOptions{
		Tty: true,
	})
	if err != nil {
		return err
	}
	defer resp.Close()

	sizeChan := make(chan DockerTermSize)
	go func() {
		for size := range sizeChan {
			cli.ContainerExecResize(ctx, execID.ID, container.ResizeOptions{
				Width:  size.Width,
				Height: size.Height,
			})
		}
	}()

	return fun(resp.Conn, sizeChan)
}

func (d *DockerApi) ContainerLogs(id string, tailCount int, fun LogsFun) error {
	cli, err := d.Conn()
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx := context.Background()

	info, err := cli.ContainerInspect(ctx, id)
	if err != nil {
		return err
	}

	opt := container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Timestamps: true,
		Follow:     true,
	}

	if tailCount != 0 {
		opt.Tail = strconv.Itoa(tailCount)
	}

	logs, err := cli.ContainerLogs(ctx, id, opt)
	if err != nil {
		return err
	}
	defer logs.Close()

	if info.Config.Tty {
		return fun(logs)
	}

	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()
		stdcopy.StdCopy(pw, pw, logs)
	}()

	return fun(pr)
}
