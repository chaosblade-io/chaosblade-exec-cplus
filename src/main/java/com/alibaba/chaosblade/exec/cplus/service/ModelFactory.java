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

import com.alibaba.chaosblade.exec.cplus.common.Constants;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

/**
 * @author Pengfei Zhou
 */
@Component
public class ModelFactory {

    @Autowired
    private DelayModelService delayModelService;
    @Autowired
    private ReturnErrorModelService returnErrorModelService;
    @Autowired
    private VariableModifyModelService variableModifyModelService;

    public IModelService createModelService(String actionName) {
        if(Constants.DELAY_ACTION_NAME.equals(actionName)){
            return delayModelService;
        } else if(Constants.RETURN_ERROR_DATA_ACTION_NAME.equals(actionName)){
            return returnErrorModelService;
        } else if(Constants.VARIABLE_MODIFY_ACTION_NAME.equals(actionName)){
            return variableModifyModelService;
        }else {
            return null;
        }
    }
}
