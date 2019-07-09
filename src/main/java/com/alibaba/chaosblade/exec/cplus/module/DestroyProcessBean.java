package com.alibaba.chaosblade.exec.cplus.module;

/**
 * @author Pengfei Zhou
 */
public class DestroyProcessBean {
    private String suid;

    public String getSuid() {
        return suid;
    }

    public void setSuid(String suid) {
        this.suid = suid;
    }

    @Override
    public String toString() {
        return "RemoveProcessBean{" +
                "suid='" + suid + '\'' +
                '}';
    }
}
