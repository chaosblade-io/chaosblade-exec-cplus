package com.alibaba.chaosblade.exec.cplus.common;

/**
 * @author Pengfei Zhou
 */
public interface Constants {

    String TARGET_NAME = "cplus";

    String DELAY_ACTION_NAME = "delay";

    String RETURN_ERROR_DATA_ACTION_NAME = "return";

    String VARIABLE_MODIFY_ACTION_NAME = "modify";

    String File_KEYWORD = "file ";

    String SET_ENV_LD_LIBRARY_PATH = "set env LD_LIBRARY_PATH ";

    Integer SLEEP_MILLISECONDS = 1000;

    Integer CORE_POOL_SIZE = 1;

    Integer MAXIMUM_POOL_SIZE = 1;

    Long KEEP_ALIVE_TIME = 0L;
}

