# petstore

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasyapi.dev/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasyapi.dev/docs/advanced-setup/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README

<!-- Start Summary [summary] -->
## Summary

Petstore - OpenAPI 3.1: This is a sample Pet Store Server based on the OpenAPI 3.1 specification.

Some useful links:
- [OpenAPI Reference](https://www.speakeasyapi.dev/openapi)
- [The Pet Store repository](https://github.com/swagger-api/swagger-petstore)
- [The source API definition for the Pet Store](https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml)

For more information about the API: [Find out more about Swagger](http://swagger.io)
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents

* [SDK Installation](#sdk-installation)
* [IDE Support](#ide-support)
* [SDK Example Usage](#sdk-example-usage)
* [Available Resources and Operations](#available-resources-and-operations)
* [File uploads](#file-uploads)
* [Retries](#retries)
* [Error Handling](#error-handling)
* [Server Selection](#server-selection)
* [Custom HTTP Client](#custom-http-client)
* [Authentication](#authentication)
* [Debugging](#debugging)
<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

The SDK can be installed with either *pip* or *poetry* package managers.

### PIP

*PIP* is the default package installer for Python, enabling easy installation and management of packages from PyPI via the command line.

```bash
pip install git+<UNSET>.git
```

### Poetry

*Poetry* is a modern tool that simplifies dependency management and package publishing by using a single `pyproject.toml` file to handle project metadata and dependencies.

```bash
poetry add git+<UNSET>.git
```
<!-- End SDK Installation [installation] -->

<!-- Start IDE Support [idesupport] -->
## IDE Support

### PyCharm

Generally, the SDK will work well with most IDEs out of the box. However, when using PyCharm, you can enjoy much better integration with Pydantic by installing an additional plugin.

- [PyCharm Pydantic Plugin](https://docs.pydantic.dev/latest/integrations/pycharm/)
<!-- End IDE Support [idesupport] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```python
# Synchronous Example
from petstore import Petstore

with Petstore(
    api_key="<YOUR_API_KEY_HERE>",
) as s:
    res = s.pet.update_pet(request={
        "name": "doggie",
        "photo_urls": [
            "<value>",
            "<value>",
        ],
        "id": 10,
        "category": {
            "id": 1,
            "name": "Dogs",
        },
    })

    if res is not None:
        # handle response
        pass
```

</br>

The same SDK client can also be used to make asychronous requests by importing asyncio.
```python
# Asynchronous Example
import asyncio
from petstore import Petstore

async def main():
    async with Petstore(
        api_key="<YOUR_API_KEY_HERE>",
    ) as s:
        res = await s.pet.update_pet_async(request={
            "name": "doggie",
            "photo_urls": [
                "<value>",
                "<value>",
            ],
            "id": 10,
            "category": {
                "id": 1,
                "name": "Dogs",
            },
        })

        if res is not None:
            # handle response
            pass

asyncio.run(main())
```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [pet](docs/sdks/petsdk/README.md)

* [update_pet](docs/sdks/petsdk/README.md#update_pet) - Update an existing pet
* [add_pet](docs/sdks/petsdk/README.md#add_pet) - Add a new pet to the store
* [find_pets_by_status](docs/sdks/petsdk/README.md#find_pets_by_status) - Finds Pets by status
* [find_pets_by_tags](docs/sdks/petsdk/README.md#find_pets_by_tags) - Finds Pets by tags
* [get_pet_by_id](docs/sdks/petsdk/README.md#get_pet_by_id) - Find pet by ID
* [delete_pet](docs/sdks/petsdk/README.md#delete_pet) - Deletes a pet
* [upload_file](docs/sdks/petsdk/README.md#upload_file) - uploads an image


### [store](docs/sdks/store/README.md)

* [get_inventory](docs/sdks/store/README.md#get_inventory) - Returns pet inventories by status
* [place_order](docs/sdks/store/README.md#place_order) - Place an order for a pet
* [get_order_by_id](docs/sdks/store/README.md#get_order_by_id) - Find purchase order by ID
* [delete_order](docs/sdks/store/README.md#delete_order) - Delete purchase order by ID

### [user](docs/sdks/usersdk/README.md)

* [create_user](docs/sdks/usersdk/README.md#create_user) - Create user
* [create_users_with_list_input](docs/sdks/usersdk/README.md#create_users_with_list_input) - Creates list of users with given input array
* [login_user](docs/sdks/usersdk/README.md#login_user) - Logs user into the system
* [logout_user](docs/sdks/usersdk/README.md#logout_user) - Logs out current logged in user session
* [get_user_by_name](docs/sdks/usersdk/README.md#get_user_by_name) - Get user by user name
* [update_user](docs/sdks/usersdk/README.md#update_user) - Update user
* [delete_user](docs/sdks/usersdk/README.md#delete_user) - Delete user

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or raise an exception.

By default, an API error will raise a models.SDKError exception, which has the following properties:

| Property        | Type             | Description           |
|-----------------|------------------|-----------------------|
| `.status_code`  | *int*            | The HTTP status code  |
| `.message`      | *str*            | The error message     |
| `.raw_response` | *httpx.Response* | The raw HTTP response |
| `.body`         | *str*            | The response content  |

When custom error responses are specified for an operation, the SDK may also raise their associated exceptions. You can refer to respective *Errors* tables in SDK docs for more details on possible exception types for each operation. For example, the `update_pet_async` method may raise the following exceptions:

| Error Type                  | Status Code | Content Type     |
| --------------------------- | ----------- | ---------------- |
| models.APIErrorInvalidInput | 400         | application/json |
| models.APIErrorUnauthorized | 401         | application/json |
| models.APIErrorNotFound     | 404         | application/json |
| models.SDKError             | 4XX, 5XX    | \*/\*            |

### Example

```python
from petstore import Petstore, models

with Petstore(
    api_key="<YOUR_API_KEY_HERE>",
) as s:
    res = None
    try:
        res = s.pet.update_pet(request={
            "name": "doggie",
            "photo_urls": [
                "<value>",
                "<value>",
            ],
            "id": 10,
            "category": {
                "id": 1,
                "name": "Dogs",
            },
        })

        if res is not None:
            # handle response
            pass

    except models.APIErrorInvalidInput as e:
        # handle e.data: models.APIErrorInvalidInputData
        raise(e)
    except models.APIErrorUnauthorized as e:
        # handle e.data: models.APIErrorUnauthorizedData
        raise(e)
    except models.APIErrorNotFound as e:
        # handle e.data: models.APIErrorNotFoundData
        raise(e)
    except models.SDKError as e:
        # handle exception
        raise(e)
```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally by passing a server index to the `server_idx: int` optional parameter when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| #   | Server                              | Variables                               | Default values |
| --- | ----------------------------------- | --------------------------------------- | -------------- |
| 0   | `http://localhost:18080`            |                                         |                |
| 1   | `https://{environment}.petstore.io` | `environment: models.ServerEnvironment` | `"prod"`       |

If the selected server has variables, you may override their default values through the additional parameters made available in the SDK constructor.

#### Example

```python
from petstore import Petstore

with Petstore(
    server_idx=1,
    api_key="<YOUR_API_KEY_HERE>",
) as s:
    res = s.pet.update_pet(request={
        "name": "doggie",
        "photo_urls": [
            "<value>",
            "<value>",
        ],
        "id": 10,
        "category": {
            "id": 1,
            "name": "Dogs",
        },
    })

    if res is not None:
        # handle response
        pass

```

### Override Server URL Per-Client

The default server can also be overridden globally by passing a URL to the `server_url: str` optional parameter when initializing the SDK client instance. For example:
```python
from petstore import Petstore

with Petstore(
    server_url="http://localhost:18080",
    api_key="<YOUR_API_KEY_HERE>",
) as s:
    res = s.pet.update_pet(request={
        "name": "doggie",
        "photo_urls": [
            "<value>",
            "<value>",
        ],
        "id": 10,
        "category": {
            "id": 1,
            "name": "Dogs",
        },
    })

    if res is not None:
        # handle response
        pass

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Python SDK makes API calls using the [httpx](https://www.python-httpx.org/) HTTP library.  In order to provide a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration, you can initialize the SDK client with your own HTTP client instance.
Depending on whether you are using the sync or async version of the SDK, you can pass an instance of `HttpClient` or `AsyncHttpClient` respectively, which are Protocol's ensuring that the client has the necessary methods to make API calls.
This allows you to wrap the client with your own custom logic, such as adding custom headers, logging, or error handling, or you can just pass an instance of `httpx.Client` or `httpx.AsyncClient` directly.

For example, you could specify a header for every request that this sdk makes as follows:
```python
from petstore import Petstore
import httpx

http_client = httpx.Client(headers={"x-custom-header": "someValue"})
s = Petstore(client=http_client)
```

or you could wrap the client with your own custom logic:
```python
from petstore import Petstore
from petstore.httpclient import AsyncHttpClient
import httpx

class CustomClient(AsyncHttpClient):
    client: AsyncHttpClient

    def __init__(self, client: AsyncHttpClient):
        self.client = client

    async def send(
        self,
        request: httpx.Request,
        *,
        stream: bool = False,
        auth: Union[
            httpx._types.AuthTypes, httpx._client.UseClientDefault, None
        ] = httpx.USE_CLIENT_DEFAULT,
        follow_redirects: Union[
            bool, httpx._client.UseClientDefault
        ] = httpx.USE_CLIENT_DEFAULT,
    ) -> httpx.Response:
        request.headers["Client-Level-Header"] = "added by client"

        return await self.client.send(
            request, stream=stream, auth=auth, follow_redirects=follow_redirects
        )

    def build_request(
        self,
        method: str,
        url: httpx._types.URLTypes,
        *,
        content: Optional[httpx._types.RequestContent] = None,
        data: Optional[httpx._types.RequestData] = None,
        files: Optional[httpx._types.RequestFiles] = None,
        json: Optional[Any] = None,
        params: Optional[httpx._types.QueryParamTypes] = None,
        headers: Optional[httpx._types.HeaderTypes] = None,
        cookies: Optional[httpx._types.CookieTypes] = None,
        timeout: Union[
            httpx._types.TimeoutTypes, httpx._client.UseClientDefault
        ] = httpx.USE_CLIENT_DEFAULT,
        extensions: Optional[httpx._types.RequestExtensions] = None,
    ) -> httpx.Request:
        return self.client.build_request(
            method,
            url,
            content=content,
            data=data,
            files=files,
            json=json,
            params=params,
            headers=headers,
            cookies=cookies,
            timeout=timeout,
            extensions=extensions,
        )

s = Petstore(async_client=CustomClient(httpx.AsyncClient()))
```
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name      | Type   | Scheme  |
| --------- | ------ | ------- |
| `api_key` | apiKey | API key |

To authenticate with the API the `api_key` parameter must be set when initializing the SDK client instance. For example:
```python
from petstore import Petstore

with Petstore(
    api_key="<YOUR_API_KEY_HERE>",
) as s:
    res = s.pet.update_pet(request={
        "name": "doggie",
        "photo_urls": [
            "<value>",
            "<value>",
        ],
        "id": 10,
        "category": {
            "id": 1,
            "name": "Dogs",
        },
    })

    if res is not None:
        # handle response
        pass

```
<!-- End Authentication [security] -->

<!-- Start File uploads [file-upload] -->
## File uploads

Certain SDK methods accept file objects as part of a request body or multi-part request. It is possible and typically recommended to upload files as a stream rather than reading the entire contents into memory. This avoids excessive memory consumption and potentially crashing with out-of-memory errors when working with very large files. The following example demonstrates how to attach a file stream to a request.

> [!TIP]
>
> For endpoints that handle file uploads bytes arrays can also be used. However, using streams is recommended for large files.
>

```python
from petstore import Petstore

with Petstore(
    api_key="<YOUR_API_KEY_HERE>",
) as s:
    res = s.pet.upload_file(pet_id=565380)

    if res is not None:
        # handle response
        pass

```
<!-- End File uploads [file-upload] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `RetryConfig` object to the call:
```python
from petstore import Petstore
from petstore.utils import BackoffStrategy, RetryConfig

with Petstore(
    api_key="<YOUR_API_KEY_HERE>",
) as s:
    res = s.pet.update_pet(request={
        "name": "doggie",
        "photo_urls": [
            "<value>",
            "<value>",
        ],
        "id": 10,
        "category": {
            "id": 1,
            "name": "Dogs",
        },
    },
        RetryConfig("backoff", BackoffStrategy(1, 50, 1.1, 100), False))

    if res is not None:
        # handle response
        pass

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `retry_config` optional parameter when initializing the SDK:
```python
from petstore import Petstore
from petstore.utils import BackoffStrategy, RetryConfig

with Petstore(
    retry_config=RetryConfig("backoff", BackoffStrategy(1, 50, 1.1, 100), False),
    api_key="<YOUR_API_KEY_HERE>",
) as s:
    res = s.pet.update_pet(request={
        "name": "doggie",
        "photo_urls": [
            "<value>",
            "<value>",
        ],
        "id": 10,
        "category": {
            "id": 1,
            "name": "Dogs",
        },
    })

    if res is not None:
        # handle response
        pass

```
<!-- End Retries [retries] -->

<!-- Start Debugging [debug] -->
## Debugging

You can setup your SDK to emit debug logs for SDK requests and responses.

You can pass your own logger class directly into your SDK.
```python
from petstore import Petstore
import logging

logging.basicConfig(level=logging.DEBUG)
s = Petstore(debug_logger=logging.getLogger("petstore"))
```
<!-- End Debugging [debug] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
