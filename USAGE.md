<!-- Start SDK Example Usage [usage] -->
```python
# Synchronous Example
from petstore import Petstore

with Petstore(
    api_key="<YOUR_API_KEY_HERE>",
) as petstore:

    res = petstore.pet.update_pet(request={
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

    assert res is not None

    # Handle response
    print(res)
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
    ) as petstore:

        res = await petstore.pet.update_pet_async(request={
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

        assert res is not None

        # Handle response
        print(res)

asyncio.run(main())
```
<!-- End SDK Example Usage [usage] -->