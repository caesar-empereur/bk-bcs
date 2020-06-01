/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package app

import (
	"os"

	"bk-bcs/bcs-common/common/blog"
	"bk-bcs/bcs-common/common/static"
	"bk-bcs/bcs-services/bcs-cpuset-device/app/options"
	"bk-bcs/bcs-services/bcs-cpuset-device/config"
	"bk-bcs/bcs-services/bcs-cpuset-device/cpuset-device"
)

func Run(op *options.Option) error {

	conf := config.NewConfig()
	setConfig(conf, op)

	controller := cpuset_device.NewCpusetDevicePlugin(conf)
	err := controller.Start()
	if err != nil {
		blog.Errorf("CpusetDevicePlugin Start failed: %s", err.Error())
		os.Exit(1)
	}

	blog.Info("CpusetDevicePlugin server ... ")
	return nil
}

func setConfig(conf *config.Config, op *options.Option) {
	conf.DockerSocket = op.DockerSock
	conf.PluginSocketDir = op.PluginSocketDir
	conf.BcsZk = op.BCSZk
	conf.Engine = op.Engine
	conf.ClusterId = op.ClusterId
	conf.NodeIp = op.Address

	//client cert directoty
	if op.CertConfig.ClientCertFile != "" && op.CertConfig.CAFile != "" &&
		op.CertConfig.ClientKeyFile != "" {

		conf.ClientCert.CertFile = op.CertConfig.ClientCertFile
		conf.ClientCert.KeyFile = op.CertConfig.ClientKeyFile
		conf.ClientCert.CAFile = op.CertConfig.CAFile
		conf.ClientCert.IsSSL = true
		conf.ClientCert.CertPasswd = static.ClientCertPwd
	}
}