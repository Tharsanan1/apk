Feature: API Deployment and invocation
  Scenario: Deploying an API and basic http method invocations
    Given The system is ready
    And I have a valid subscription
    When I use the APK Conf file "artifacts/apk-confs/employees_conf.yaml"
    And the definition file "artifacts/definitions/employees_api.json"
    And make the API deployment request
    Then the response status code should be 200
    Then I set headers
      |Authorization|bearer ${accessToken}|
    And I send "GET" request to "https://default.gw.wso2.com:9095/test/3.14/employee/" with body ""
    And I eventually receive 200 response code, not accepting
      |429|
    And I send "POST" request to "https://default.gw.wso2.com:9095/test/3.14/employee/" with body ""
    And I receive 200 response code
    And I send "POST" request to "https://default.gw.wso2.com:9095/test/3.14/test/" with body ""
    And I receive 404 response code
    And I send "PUT" request to "https://default.gw.wso2.com:9095/test/3.14/employee/12" with body ""
    And I receive 200 response code
    And I send "DELETE" request to "https://default.gw.wso2.com:9095/test/3.14/employee/12" with body ""
    And I receive 200 response code

  Scenario: Deploying an API with default version
    Given The system is ready
    And I have a valid subscription
    When I use the APK Conf file "artifacts/apk-confs/default_version_conf.yaml"
    And the definition file "artifacts/definitions/employees_api.json"
    And make the API deployment request
    Then the response status code should be 200
    Then I set headers
      |Authorization|bearer ${accessToken}|
    And I send "GET" request to "https://default.gw.wso2.com:9095/test-default/3.14/employee/" with body ""
    And I eventually receive 200 response code, not accepting
      |429|
    And I send "POST" request to "https://default.gw.wso2.com:9095/test-default/3.14/employee/" with body ""
    And I receive 200 response code
    And I send "POST" request to "https://default.gw.wso2.com:9095/test-default/3.14/test/" with body ""
    And I receive 404 response code
    And I send "PUT" request to "https://default.gw.wso2.com:9095/test-default/3.14/employee/12" with body ""
    And I receive 200 response code
    And I send "DELETE" request to "https://default.gw.wso2.com:9095/test-default/3.14/employee/12" with body ""
    And I receive 200 response code

    And I send "GET" request to "https://default.gw.wso2.com:9095/test-default/employee/" with body ""
    And I eventually receive 200 response code, not accepting
      |429|
    And I send "POST" request to "https://default.gw.wso2.com:9095/test-default/employee/" with body ""
    And I receive 200 response code
    And I send "POST" request to "https://default.gw.wso2.com:9095/test-default/test/" with body ""
    And I receive 404 response code
    And I send "PUT" request to "https://default.gw.wso2.com:9095/test-default/employee/12" with body ""
    And I receive 200 response code
    And I send "DELETE" request to "https://default.gw.wso2.com:9095/test-default/employee/12" with body ""
    And I receive 200 response code

  Scenario Outline: Undeploy API
    Given The system is ready
    And I have a valid subscription
    When I undeploy the API whose ID is "<apiID>"
    Then the response status code should be <expectedStatusCode>

    Examples:
      | apiID                 | expectedStatusCode  |
      | emp-api-test          | 202                 |
      | default-version-api-test      | 202                 |
