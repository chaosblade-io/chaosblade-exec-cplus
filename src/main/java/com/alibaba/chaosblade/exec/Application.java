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

package com.alibaba.chaosblade.exec;

import com.alibaba.chaosblade.exec.cplus.common.Response;
import com.alibaba.chaosblade.exec.cplus.common.Result;
import com.alibaba.chaosblade.exec.cplus.utils.ExecShellUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.core.env.Environment;

/**
 * @author Pengfei Zhou
 */
@SpringBootApplication
public class Application
{
    private static final Logger logger = LoggerFactory.getLogger(Application.class);

    public static void main( String[] args )
    {
        ConfigurableApplicationContext context = SpringApplication.run(Application.class, args);
        Environment environment = context.getBean(Environment.class);
        String location = environment.getProperty("script.location");
        String initializationShellName = environment.getProperty("script.initialization");
        Result initializationResult = ExecShellUtils.execShell(location + initializationShellName);
        if (!initializationResult.isSuccess()) {
            logger.info("Fail to install" + "Error code: " + Response.Code.SERVER_ERROR + " error message: " + initializationResult.getErr());
        } else {
            logger.info("Succeed to install");
        }
    }
}