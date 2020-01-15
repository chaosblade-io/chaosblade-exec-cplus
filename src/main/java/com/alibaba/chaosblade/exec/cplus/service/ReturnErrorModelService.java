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

import com.alibaba.chaosblade.exec.cplus.module.RequestParamsBean;
import com.alibaba.chaosblade.exec.cplus.utils.ExecShellUtils;
import com.alibaba.chaosblade.exec.cplus.common.Result;
import com.alibaba.chaosblade.exec.cplus.utils.StringUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

/**
 * @author Pengfei Zhou
 */
@Component
public class ReturnErrorModelService implements IModelService {

    private String strScriptLocation;
    private String strScriptReturnErrorFileName;
    private String strScriptReturnErrorAttachFileName;

    @Autowired
    private CheckStatusService checkStatusService;

    @Override
    public Result validateFlag(RequestParamsBean requestParamsBean) {
        String returnValue = requestParamsBean.getReturnValue();
        if (StringUtil.isBlank(returnValue)) {
            return Result.fail("less necessary returnValue value");
        }
        return Result.success();
    }

    @Override
    public Result handleInjection(RequestParamsBean requestParamsBean) {
        String pid = checkStatusService.getProcessIdByProcessName(requestParamsBean.getProcessName());

        if (StringUtil.isBlank(pid)){
            return ExecShellUtils.execShell(strScriptLocation + strScriptReturnErrorFileName, requestParamsBean.getFileLocateAndName(),requestParamsBean.getForkMode()
                    ,requestParamsBean.getLibLoad(),requestParamsBean.getBreakLine(),requestParamsBean.getReturnValue(),requestParamsBean.getInitParams());
        } else {
            return ExecShellUtils.execShell(strScriptLocation + strScriptReturnErrorAttachFileName, pid, requestParamsBean.getForkMode(), "", ""
                    , requestParamsBean.getBreakLine(), requestParamsBean.getReturnValue(), requestParamsBean.getInitParams());
        }
    }

    @Autowired
    public void ReturnErrorModelService(@Value("${script.location}") String location,
                                  @Value("${script.return.error.file.name}") String name,
                                        @Value("${script.return.error.attach.file.name}") String nameAttach) {
        strScriptLocation = location;
        strScriptReturnErrorFileName = name;
        strScriptReturnErrorAttachFileName = nameAttach;
    }
}
