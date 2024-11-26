<!-- Start SDK Example Usage [usage] -->
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