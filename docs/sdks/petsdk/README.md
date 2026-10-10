# Pet

## Overview

Everything about your Pets

Find out more
<http://swagger.io>

### Available Operations

* [update_pet](#update_pet) - Update an existing pet
* [add_pet](#add_pet) - Add a new pet to the store
* [find_pets_by_status](#find_pets_by_status) - Finds Pets by status
* [find_pets_by_tags](#find_pets_by_tags) - Finds Pets by tags
* [get_pet_by_id](#get_pet_by_id) - Find pet by ID
* [delete_pet](#delete_pet) - Deletes a pet
* [upload_file](#upload_file) - uploads an image

## update_pet

Update an existing pet by Id

### Example Usage

<!-- UsageSnippet language="python" operationID="updatePet" method="put" path="/pet" -->
```python
from petstore import Petstore


with Petstore(
    api_key="<YOUR_API_KEY_HERE>",
) as p_client:

    res = p_client.pet.update_pet(request={
        "id": 10,
        "name": "doggie",
        "category": {
            "id": 1,
            "name": "Dogs",
        },
        "photo_urls": [
            "<value 1>",
        ],
    })

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `request`                                                           | [models.Pet](../../models/pet.md)                                   | :heavy_check_mark:                                                  | The request object to use for the request.                          |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.Pet](../../models/pet.md)**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| models.APIErrorInvalidInput | 400                         | application/json            |
| models.APIErrorUnauthorized | 401                         | application/json            |
| models.APIErrorNotFound     | 404                         | application/json            |
| models.SDKError             | 4XX, 5XX                    | \*/\*                       |

## add_pet

Add a new pet to the store

### Example Usage

<!-- UsageSnippet language="python" operationID="addPet" method="post" path="/pet" -->
```python
from petstore import Petstore


with Petstore(
    api_key="<YOUR_API_KEY_HERE>",
) as p_client:

    res = p_client.pet.add_pet(request={
        "id": 10,
        "name": "doggie",
        "category": {
            "id": 1,
            "name": "Dogs",
        },
        "photo_urls": [
            "<value 1>",
            "<value 2>",
            "<value 3>",
        ],
    })

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `request`                                                           | [models.Pet](../../models/pet.md)                                   | :heavy_check_mark:                                                  | The request object to use for the request.                          |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.Pet](../../models/pet.md)**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| models.SDKError | 4XX, 5XX        | \*/\*           |

## find_pets_by_status

Multiple status values can be provided with comma separated strings

### Example Usage

<!-- UsageSnippet language="python" operationID="findPetsByStatus" method="get" path="/pet/findByStatus" -->
```python
import petstore
from petstore import Petstore


with Petstore(
    api_key="<YOUR_API_KEY_HERE>",
) as p_client:

    res = p_client.pet.find_pets_by_status(status=petstore.QueryParamStatus.AVAILABLE)

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                             | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `status`                                                              | [Optional[models.QueryParamStatus]](../../models/queryparamstatus.md) | :heavy_minus_sign:                                                    | Status values that need to be considered for filter                   |
| `retries`                                                             | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)      | :heavy_minus_sign:                                                    | Configuration to override the default retry behavior of the client.   |

### Response

**[List[models.Pet]](../../models/.md)**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| models.APIErrorInvalidInput | 400                         | application/json            |
| models.APIErrorUnauthorized | 401                         | application/json            |
| models.APIErrorNotFound     | 404                         | application/json            |
| models.SDKError             | 4XX, 5XX                    | \*/\*                       |

## find_pets_by_tags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

### Example Usage

<!-- UsageSnippet language="python" operationID="findPetsByTags" method="get" path="/pet/findByTags" -->
```python
from petstore import Petstore


with Petstore(
    api_key="<YOUR_API_KEY_HERE>",
) as p_client:

    res = p_client.pet.find_pets_by_tags()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `tags`                                                              | List[*str*]                                                         | :heavy_minus_sign:                                                  | Tags to filter by                                                   |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[List[models.Pet]](../../models/.md)**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| models.APIErrorInvalidInput | 400                         | application/json            |
| models.APIErrorUnauthorized | 401                         | application/json            |
| models.APIErrorNotFound     | 404                         | application/json            |
| models.SDKError             | 4XX, 5XX                    | \*/\*                       |

## get_pet_by_id

Returns a single pet

### Example Usage

<!-- UsageSnippet language="python" operationID="getPetById" method="get" path="/pet/{petId}" -->
```python
from petstore import Petstore


with Petstore(
    api_key="<YOUR_API_KEY_HERE>",
) as p_client:

    res = p_client.pet.get_pet_by_id(pet_id=311674)

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `pet_id`                                                            | *int*                                                               | :heavy_check_mark:                                                  | ID of pet to return                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.Pet](../../models/pet.md)**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| models.APIErrorInvalidInput | 400                         | application/json            |
| models.APIErrorUnauthorized | 401                         | application/json            |
| models.APIErrorNotFound     | 404                         | application/json            |
| models.SDKError             | 4XX, 5XX                    | \*/\*                       |

## delete_pet

Deletes a pet

### Example Usage

<!-- UsageSnippet language="python" operationID="deletePet" method="delete" path="/pet/{petId}" -->
```python
from petstore import Petstore


with Petstore(
    api_key="<YOUR_API_KEY_HERE>",
) as p_client:

    res = p_client.pet.delete_pet(pet_id=818965)

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `pet_id`                                                            | *int*                                                               | :heavy_check_mark:                                                  | Pet id to delete                                                    |
| `api_key`                                                           | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.Pet](../../models/pet.md)**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| models.APIErrorInvalidInput | 400                         | application/json            |
| models.APIErrorUnauthorized | 401                         | application/json            |
| models.APIErrorNotFound     | 404                         | application/json            |
| models.SDKError             | 4XX, 5XX                    | \*/\*                       |

## upload_file

uploads an image

### Example Usage

<!-- UsageSnippet language="python" operationID="uploadFile" method="post" path="/pet/{petId}/uploadImage" -->
```python
from petstore import Petstore


with Petstore(
    api_key="<YOUR_API_KEY_HERE>",
) as p_client:

    res = p_client.pet.upload_file(pet_id=150516)

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `pet_id`                                                            | *int*                                                               | :heavy_check_mark:                                                  | ID of pet to update                                                 |
| `additional_metadata`                                               | *Optional[str]*                                                     | :heavy_minus_sign:                                                  | Additional Metadata                                                 |
| `request_body`                                                      | *Optional[Union[bytes, IO[bytes], io.IOBase]]*                      | :heavy_minus_sign:                                                  | N/A                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[models.APIResponse](../../models/apiresponse.md)**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| models.SDKError | 4XX, 5XX        | \*/\*           |