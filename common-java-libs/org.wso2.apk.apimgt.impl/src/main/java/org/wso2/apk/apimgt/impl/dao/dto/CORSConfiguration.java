/*
 * Copyright (c) 2022, WSO2 LLC. (http://www.wso2.com).
 *
 * WSO2 LLC. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package org.wso2.apk.apimgt.impl.dao.dto;

import java.util.List;

/**
 * API CORS Configuration
 */
public class CORSConfiguration {

    private boolean corsConfigurationEnabled;
    private List<String> accessControlAllowOrigins;
    private boolean accessControlAllowCredentials;
    private List<String> accessControlAllowHeaders;
    private List<String> accessControlAllowMethods;

    public boolean isCorsConfigurationEnabled() {
        return corsConfigurationEnabled;
    }

    public void setCorsConfigurationEnabled(boolean corsConfigurationEnabled) {
        this.corsConfigurationEnabled = corsConfigurationEnabled;
    }

    public List<String> getAccessControlAllowOrigins() {
        return accessControlAllowOrigins;
    }

    public void setAccessControlAllowOrigins(List<String> accessControlAllowOrigins) {
        this.accessControlAllowOrigins = accessControlAllowOrigins;
    }

    public boolean isAccessControlAllowCredentials() {
        return accessControlAllowCredentials;
    }

    public void setAccessControlAllowCredentials(boolean accessControlAllowCredentials) {
        this.accessControlAllowCredentials = accessControlAllowCredentials;
    }

    public List<String> getAccessControlAllowHeaders() {
        return accessControlAllowHeaders;
    }

    public void setAccessControlAllowHeaders(List<String> accessControlAllowHeaders) {
        this.accessControlAllowHeaders = accessControlAllowHeaders;
    }

    public List<String> getAccessControlAllowMethods() {
        return accessControlAllowMethods;
    }

    public void setAccessControlAllowMethods(List<String> accessControlAllowMethods) {
        this.accessControlAllowMethods = accessControlAllowMethods;
    }
    
    @Override
    public String toString() {
        return "CORSConfiguration [corsConfigurationEnabled=" + corsConfigurationEnabled
                + ", accessControlAllowOrigins=" + accessControlAllowOrigins + ", accessControlAllowCredentials="
                + accessControlAllowCredentials + ", accessControlAllowHeaders=" + accessControlAllowHeaders
                + ", accessControlAllowMethods=" + accessControlAllowMethods + "]";
    }
}
