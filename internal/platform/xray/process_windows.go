package xray

import (
	"os/exec"
	"strconv"
	"strings"
)

func (p *WindowsProcessManager) detectExternalProcess() Status {

	out, err := exec.Command(
		"tasklist",
		"/FI",
		"IMAGENAME eq xray.exe",
		"/FO",
		"CSV",
	).Output()

	if err != nil {
		return Status{
			Running: false,
			PID:     0,
		}
	}

	lines := strings.Split(
		string(out),
		"\n",
	)

	for _, line := range lines {

		if !strings.Contains(
			strings.ToLower(line),
			"xray.exe",
		) {
			continue
		}

		fields := strings.Split(
			line,
			",",
		)

		if len(fields) < 2 {
			continue
		}

		pidString := strings.Trim(
			fields[1],
			"\"",
		)

		pid, err := strconv.Atoi(pidString)

		if err != nil {
			continue
		}

		return Status{
			Running: true,
			PID:     pid,
		}
	}

	return Status{
		Running: false,
		PID:     0,
	}
}
