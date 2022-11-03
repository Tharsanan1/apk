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

package org.wso2.apk.apimgt.api.model;

import java.util.ArrayList;
import java.util.List;

public class ConfigurationDto {

    private String name;
    private String label;
    private String type;
    private String tooltip;
    private Object defaultValue;
    private boolean required;
    private boolean mask;
    private List values = new ArrayList<>();
    private boolean multiple;

    public String getName() {

        return name;
    }

    public void setName(String name) {

        this.name = name;
    }

    public String getLabel() {

        return label;
    }

    public void setLabel(String label) {

        this.label = label;
    }

    public String getType() {

        return type;
    }

    public void setType(String type) {

        this.type = type;
    }

    public String getTooltip() {

        return tooltip;
    }

    public void setTooltip(String tooltip) {

        this.tooltip = tooltip;
    }

    public Object getDefaultValue() {

        return defaultValue;
    }

    public void setDefaultValue(String defaultValue) {

        this.defaultValue = defaultValue;
    }

    public boolean isRequired() {

        return required;
    }

    public void setRequired(boolean required) {

        this.required = required;
    }

    public boolean isMask() {

        return mask;
    }

    public void setMask(boolean mask) {

        this.mask = mask;
    }

    public List<Object> getValues() {

        return values;
    }

    public void setValues(List<String> values) {

        this.values = values;
    }

    public boolean isMultiple() {

        return multiple;
    }

    public void setMultiple(boolean multiple) {

        this.multiple = multiple;
    }

    public void addValue(String value) {

        this.values.add(value);
    }

    public ConfigurationDto(String name, String label, String type, String tooltip, Object defaultValue,
                            boolean required,
                            boolean mask, List values, boolean multiple) {

        this.name = name;
        this.label = label;
        this.type = type;
        this.tooltip = tooltip;
        this.defaultValue = defaultValue;
        this.required = required;
        this.mask = mask;
        this.values = values;
        this.multiple = multiple;
    }
    public ConfigurationDto(String name, String label, String type, String tooltip, String defaultValue,
                            boolean required,
                            boolean mask, List values, boolean multiple) {

        this.name = name;
        this.label = label;
        this.type = type;
        this.tooltip = tooltip;
        this.defaultValue = defaultValue;
        this.required = required;
        this.mask = mask;
        this.values = values;
        this.multiple = multiple;
    }
}
