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

package com.alibaba.chaosblade.exec.cplus.service;

import com.alibaba.chaosblade.exec.cplus.common.Result;
import com.alibaba.chaosblade.exec.cplus.utils.ExecShellUtils;
import com.alibaba.chaosblade.exec.cplus.utils.StringUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

/**
 * @author Pengfei Zhou
 */
@Component
public class CheckStatusService {
    private static String strScriptDelayLocation;
    private static String strScriptCheckProcessId;

    public String getProcessIdByProcessName(String processName){
        Result result = ExecShellUtils.execShell(strScriptDelayLocation + strScriptCheckProcessId, processName);
        if(result != null && result.isSuccess() && !StringUtil.isBlank(result.getComment())){
            return result.getComment();
        }
        return "";
    }

    @Autowired
    public void CheckStatusService(@Value("${script.location}") String location,
                                  @Value("${script.check.process.id}") String name) {
        strScriptDelayLocation = location;
        strScriptCheckProcessId = name;
    }
}
