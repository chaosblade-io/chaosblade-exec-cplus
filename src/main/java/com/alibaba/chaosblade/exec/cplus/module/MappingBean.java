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
