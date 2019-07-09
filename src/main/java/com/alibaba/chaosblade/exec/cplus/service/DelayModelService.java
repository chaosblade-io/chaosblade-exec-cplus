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
