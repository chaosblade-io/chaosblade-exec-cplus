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
public class VariableModifyModelService implements IModelService {

    private String strScriptLocation;
    private String strScriptVariableModifyFileName;
    private String strScriptVariableModifyAttachFileName;

    @Autowired
    private CheckStatusService checkStatusService;

    @Override
    public Result validateFlag(RequestParamsBean requestParamsBean) {
        String varaibleName = requestParamsBean.getVaraibleName();
        if (StringUtil.isBlank(varaibleName)) {
            return Result.fail("less necessary varaibleName value");
        }

        String varaibleValue = requestParamsBean.getVaraibleValue();
        if (StringUtil.isBlank(varaibleValue)) {
            return Result.fail("less necessary varaibleValue value");
        }
        return Result.success();
    }

    @Override
    public Result handleInjection(RequestParamsBean requestParamsBean) {
        String pid = checkStatusService.getProcessIdByProcessName(requestParamsBean.getProcessName());

        if (StringUtil.isBlank(pid)){
            return ExecShellUtils.execShell(strScriptLocation + strScriptVariableModifyFileName, requestParamsBean.getFileLocateAndName(),requestParamsBean.getForkMode()
                    ,requestParamsBean.getLibLoad(),requestParamsBean.getBreakLine(),requestParamsBean.getVaraibleName(),requestParamsBean.getVaraibleValue(),requestParamsBean.getInitParams());
        } else {
            return ExecShellUtils.execShell(strScriptLocation + strScriptVariableModifyAttachFileName, pid, requestParamsBean.getForkMode(), "", ""
                    , requestParamsBean.getBreakLine(), requestParamsBean.getVaraibleName(), requestParamsBean.getVaraibleValue(), requestParamsBean.getInitParams());
        }
    }

    @Autowired
    public void VariableModifyModelService(@Value("${script.location}") String location,
                                        @Value("${script.variable.modify.file.name}") String name,
                                           @Value("${script.return.error.attach.file.name}") String nameAttach) {
        strScriptLocation = location;
        strScriptVariableModifyFileName = name;
        strScriptVariableModifyAttachFileName = nameAttach;
    }
}
