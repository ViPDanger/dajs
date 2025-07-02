# Dungeons And JSON

Сервис позволяющий работать с сохранениями Helpmate2 через REST API

Основные конфиги с настройками приложения лежат в
cmd/api/config.ini
cmd/worker/config.ini

Список запросов: 
  |метод   |  путь                       |  body                       |  response                                                    |     Описание   				    |
  |--------|-----------------------------|-----------------------------|--------------------------------------------------------------|-----------------------------------------------------------------------------------------|
  |GET     |  /                          |                             | OK                                                           |   Проверка доступности сервера		    |				
  |AUTH       					    |								
  |POST    |  auth/register              | {"user","password"}         | {"Message"}                                                  |   Регистрация пользователя		    |						
  |POST    |  auth/login                 |  {"user","password"}        | {"access_token","access_exp","refresh_token","refresh_exp"}  |   Авторизация авториза			    |
  |POST    |  auth/refresh               |  {"refresh_token"}          |  {"access_token","access_exp"}                               |   Получение access токена из refresh токена |     
  |PROTECTED*|	     |		 					    |
  |GET     |  protected/characrer/all      |                             |  {[]CharacterDTO}					      |  Получение массива обьектов		    |
  |GET     |  protected/characrer    |   HEADER "ID"               |  {CharacterDTO}			 		      |  Получение обьекта по его ID	    |
  |POST    |  protected/characrer    |   {CharacterDTO}            |  {"Message"}			      			      |  POST нового обьекта			    |
  |PUT     |  protected/characrer    |   {CharacterDTO}            |  {"Message"}						      |  Изменение существующего обьекта (по NON default полям)						    | 
  |DELETE  |  protected/characrer |   HEADER "ID"               |  {"Message"}						      |  Удаление обьекта из системы|
  -----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

*По аналогии с protected/character так же существуют
protected/item
protected/map
protected/glossary
