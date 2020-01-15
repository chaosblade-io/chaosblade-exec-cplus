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
import com.alibaba.chaosblade.exec.cplus.utils.StringUtil;
import com.alibaba.chaosblade.exec.cplus.common.Result;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

/**
 * @author Pengfei Zhou
 */
@Component
public class DelayModelService implements IModelService {

    private String strScriptLocation;
    private String strScriptDelayFileName;
    private String strScriptDelayAttachFileName;

    @Autowired
    private CheckStatusService checkStatusService;

    @Override
    public Result validateFlag(RequestParamsBean requestParamsBean) {
        String delayDuration = requestParamsBean.getDelayDuration();
        if (StringUtil.isBlank(delayDuration)) {
            return Result.fail("less necessary delayDuration value");
        }
        return Result.success();
    }

    @Override
    public Result handleInjection(RequestParamsBean requestParamsBean) {
        String pid = checkStatusService.getProcessIdByProcessName(requestParamsBean.getProcessName());

        if (StringUtil.isBlank(pid)){
            return ExecShellUtils.execShell(strScriptLocation + strScriptDelayFileName, requestParamsBean.getFileLocateAndName(),requestParamsBean.getForkMode()
                    ,requestParamsBean.getLibLoad(),requestParamsBean.getBreakLine(),requestParamsBean.getDelayDuration(),requestParamsBean.getInitParams());
        } else {
            return ExecShellUtils.execShell(strScriptLocation + strScriptDelayAttachFileName, pid, requestParamsBean.getForkMode(), "", ""
                    , requestParamsBean.getBreakLine(), requestParamsBean.getDelayDuration(), requestParamsBean.getInitParams());
        }
    }

    @Autowired
    public void DelayModelService(@Value("${script.location}") String location,
                                  @Value("${script.delay.file.name}") String name,
                                  @Value("${script.delay.attach.file.name}") String nameAttach) {
        strScriptLocation = location;
        strScriptDelayFileName = name;
        strScriptDelayAttachFileName = nameAttach;
    }
}
