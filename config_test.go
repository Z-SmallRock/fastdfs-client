package fdfs_client

import (
	"testing"
	"fmt"
)

func TestConfig(t *testing.T) {
	config, err := newConfigByYaml([]string{"10.8.145.65:22122"},"200")
	if err != nil {
		fmt.Println(err)
		return;
	}
	fmt.Println(config.trackerAddr)
	fmt.Println(config.maxConns)
}
