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

package com.alibaba.chaosblade.exec.cplus.utils;

import com.alibaba.chaosblade.exec.cplus.common.Result;
import org.apache.commons.lang3.ArrayUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.BufferedReader;
import java.io.InputStreamReader;

/**
 * @author Pengfei Zhou
 */
public class ExecShellUtils {

    private static final Logger logger = LoggerFactory.getLogger(ExecShellUtils.class);
    /**
     * Shell script execution
     *
     * @param scriptPath
     * @param para
     */
    public static Result execShell(String scriptPath, String... para) {
        try {
            String[] cmd = new String[]{scriptPath};
            cmd = ArrayUtils.addAll(cmd, para);

            ProcessBuilder builder = new ProcessBuilder("/bin/chmod", "755", scriptPath);
            Process process = builder.start();
            process.waitFor();

            Process ps = Runtime.getRuntime().exec(cmd);
            ps.waitFor();

            BufferedReader br = new BufferedReader(new InputStreamReader(ps.getInputStream()));
            StringBuffer sb = new StringBuffer();
            String line;
            while ((line = br.readLine()) != null) {
                sb.append(line).append("\n");
            }
            logger.info("execShell return : "+ sb.toString());
            return Result.success(sb.toString());
        } catch (Exception e) {
            logger.error("execShell request exception, params: " + para, e);
            return Result.fail(e.getMessage());
        }
    }
}
