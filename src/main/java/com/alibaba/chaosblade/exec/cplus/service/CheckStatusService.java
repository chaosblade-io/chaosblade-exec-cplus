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
