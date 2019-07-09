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
