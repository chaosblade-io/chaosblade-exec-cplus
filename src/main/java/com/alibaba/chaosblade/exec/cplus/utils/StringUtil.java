package com.alibaba.chaosblade.exec.cplus.utils;

/**
 * @author Pengfei Zhou
 */
public class StringUtil {
    /**
     * Is not null or not empty
     *
     * @param value
     * @return
     */
    public static boolean isBlank(String value) {
        return value == null || (value.trim()).length() == 0;
    }
}
