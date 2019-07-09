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
public class CheckDuplicateService {
    private static String strScriptLocation;
    private static String strScriptCheckProcessDuplicate;

    public boolean isExistProcessByProcessName(String processName){
        Result result = ExecShellUtils.execShell(strScriptLocation + strScriptCheckProcessDuplicate, processName);
        if(result != null && result.isSuccess() && !StringUtil.isBlank(result.getComment())){
            return true;
        }
        return false;
    }

    @Autowired
    public void CheckDuplicateService(@Value("${script.location}") String location,
                                      @Value("${script.check.process.duplicate}") String name) {
        strScriptLocation = location;
        strScriptCheckProcessDuplicate = name;
    }
}
