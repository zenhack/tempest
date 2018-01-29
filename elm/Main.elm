import Html exposing(..)
import Html.Attributes exposing (href)
import Json.Decode exposing (list, string)
import Http

type alias Model =
  { apps : List String
  , rootHost : String
  , protocol : String
  }

type Rpc
  = ListApps

type Msg
  = AppList (List String)
  | RpcError Http.Error

sendRpc : Model -> Rpc -> Cmd Msg
sendRpc model ListApps =
    Http.get
      (model.protocol ++ "://" ++ model.rootHost ++ "/.rpc/ListApps")
      (list string)
    |>
    Http.send (\resp -> case resp of
      Ok apps -> AppList apps
      Err err -> RpcError err)

getAppDomain : Model -> String -> String
getAppDomain model subdomain =
  model.protocol ++ "://" ++ subdomain ++ "." ++ model.rootHost

update : Msg -> Model -> (Model, Cmd Msg)
update msg model = case msg of
  AppList apps ->
    ( { model | apps = apps }
    , Cmd.none
    )
  RpcError err ->
    ( { model | apps = [ "Rpc Error " ++ toString err ] }
    , Cmd.none
    )

view : Model -> Html Msg
view model =
  let viewAppLink appName =
    li [ ] [ a [ href (getAppDomain model appName) ] [ text appName ] ]
  in
    ul [ ]
      (List.map viewAppLink model.apps)

subscriptions : Model -> Sub Msg
subscriptions _ = Sub.none

init : (Model, Cmd Msg)
init =
  let model =
    { apps = [ ]
    -- TODO: extract these from the page.
    , rootHost = "local.sandstorm.io:8000"
    , protocol = "http"
    }
  in
     ( model
     , sendRpc model ListApps
     )

main = Html.program
  { init = init
  , view = view
  , update = update
  , subscriptions = subscriptions
  }
