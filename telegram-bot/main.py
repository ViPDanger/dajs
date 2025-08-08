import asyncio
from aiogram import Bot, Dispatcher,Router
from options import BOT_TOKEN
from handlers import routers
from aiohttp import web


async def main():
    app = web.Application()
    app.router.add_post("/webapp/submit", handle_webapp_submit)
    app.router.add_static("/", path="public", show_index=True)  # index.html
    runner = web.AppRunner(app)
    await runner.setup()
    site = web.TCPSite(runner, port=443, ssl_context=("cert.pem", "key.pem"))
    await site.start()
    bot = Bot(token=BOT_TOKEN)
    dp = Dispatcher()
    router = Router()
    for r in routers:
        router.include_router(r)
    dp.include_router(router)
    await dp.start_polling(bot)

if __name__ == "__main__":
    asyncio.run(main())

     