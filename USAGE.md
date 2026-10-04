<!-- Start SDK Example Usage [usage] -->
```python
# Synchronous Example
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

</br>

The same SDK client can also be used to make asynchronous requests by importing asyncio.

```python
# Asynchronous Example
import asyncio
from petstore import Petstore

async def main():

    async with Petstore(
        api_key="<YOUR_API_KEY_HERE>",
    ) as p_client:

        res = await p_client.pet.update_pet_async(request={
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

asyncio.run(main())
```
<!-- End SDK Example Usage [usage] -->