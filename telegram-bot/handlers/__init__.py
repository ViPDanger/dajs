import handlers.auth_handler as auth
import handlers.character_handler as character
import handlers.fallback_handler as fallback

routers =[
auth.router,
character.router,
fallback.router,
]

