from uuid import uuid4

from aiohttp import web


async def prefix(request):
    return web.json_response({'prefix': str(uuid4())})


app = web.Application()
app.add_routes([
    web.get('/prefix', prefix),
])


if __name__ == '__main__':
    web.run_app(app)
