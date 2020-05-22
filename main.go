/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/chaosblade-io/chaosblade-spec-go/channel"

	"github.com/chaosblade-io/chaosblade-exec-cplus/common"
	"github.com/chaosblade-io/chaosblade-exec-cplus/controller"
)

type Config struct {
	Port  int
	IP    string
	nohup bool
}

func main() {
	config := &Config{}
	flag.StringVar(&config.IP, "ip", "", "The ip bounds on the service")
	flag.IntVar(&config.Port, "port", 9525, "The port bounds one the service")
	flag.BoolVar(&config.nohup, "nohup", false, "used by internal")

	flag.Parse()
	ctx := context.WithValue(context.Background(), channel.ProcessKey, "nohup")
	pids, err := channel.NewLocalChannel().GetPidsByProcessName(common.BinName, ctx)
	if err != nil {
		log.Fatalf("query process failed, %v", err)
	}
	if len(pids) > 0 {
		log.Fatalf("process has been started, %+v", pids)
	}
	if config.nohup {
		start0(config)
		os.Exit(0)
	}
	err = start(config)
	if err != nil {
		log.Fatalf("start failed, %v", err)
	}
	log.Printf("success")
	os.Exit(0)
}

func start(config *Config) error {
	args := fmt.Sprintf("%s --nohup --port %d", path.Join(common.GetProgramPath(), common.BinName), config.Port)
	if config.IP != "" {
		args = fmt.Sprintf("%s --ip %s", args, config.IP)
	}
	cl := channel.NewLocalChannel()

	response := cl.Run(context.TODO(), "nohup", fmt.Sprintf("%s > %s 2>&1 &", args, common.GetChaosBladeLogPath()))
	if !response.Success {
		return fmt.Errorf(response.Err)
	}
	time.Sleep(time.Second)
	ctx := context.WithValue(context.Background(), channel.ProcessKey, "nohup")
	pids, err := cl.GetPidsByProcessName(common.BinName, ctx)
	if err != nil {
		return err
	}
	if len(pids) == 0 {
		return fmt.Errorf("process not found")
	}
	return nil
}

func start0(config *Config) {
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("%s:%d", config.IP, config.Port), nil)
		if err != nil {
			log.Fatalf("start chaosblade-exec-cplus failed, %v", err)
		}
	}()

	for _, c := range controller.Controllers {
		http.HandleFunc("/"+c.GetControllerName(), c.GetRequestHandler())
	}
	common.Hold()
}
