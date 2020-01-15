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

package com.alibaba.chaosblade.exec.cplus.module;

/**
 * @author Pengfei Zhou
 */
public class MappingBean {
    private volatile static MappingBean mappingBean;

    private MappingBean(){
    }

    public static MappingBean getInstance() {
        if (mappingBean == null) {
            synchronized (MappingBean.class) {
                if (mappingBean == null) {
                    mappingBean = new MappingBean();
                }
            }
        }
        return mappingBean;
    }

    private String suid;

    private String processName;

    public String getSuid() {
        return suid;
    }

    public void setSuid(String suid) {
        this.suid = suid;
    }

    public String getProcessName() {
        return processName;
    }

    public void setProcessName(String processName) {
        this.processName = processName;
    }

}
