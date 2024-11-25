<!-- Start SDK Example Usage [usage] -->
```python
# Synchronous Example
from petstore import Petstore

s = Petstore(
    api_key="<YOUR_API_KEY_HERE>",
)

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
    s = Petstore(
        api_key="<YOUR_API_KEY_HERE>",
    )
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