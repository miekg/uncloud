package docker

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/docker/docker/api/types/filters"
)

// PruneImages wakes up every hour to prune the docker images.
func (c *Controller) PruneImages(ctx context.Context) error {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	// see https://docs.docker.com/reference/cli/docker/container/prune/
	filter := filters.NewArgs(
		filters.Arg("until", "24h"),
	)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			report, err := c.client.ContainersPrune(ctx, filter)
			if err == nil {
				slog.Info("Pruned", "images deleted", len(report.ContainersDeleted), "reclaimed", fmt.Sprintf("%d B", report.SpaceReclaimed))
			}
		}
	}
}
