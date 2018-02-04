module Main exposing (..)

import Html exposing (..)
import Html.Attributes exposing (href)
import Json.Decode exposing (list, string)
import Http
import WebSocket


type alias ServerAddr =
    { rootHost : String
    , secure : Bool
    }


proto : String -> ServerAddr -> String
proto insecure addr =
    if addr.secure then
        insecure ++ "s"
    else
        insecure


urlFor : String -> String -> ServerAddr -> String
urlFor p path addr =
    proto p addr ++ "://" ++ addr.rootHost ++ "/" ++ path


type alias Model =
    { apps : List String
    , serverAddr : ServerAddr
    }


type Rpc
    = ListApps


type Msg
    = AppList (List String)
    | RpcError Http.Error
    | NewData String


sendRpc : Model -> Rpc -> Cmd Msg
sendRpc model ListApps =
    Http.get
        (urlFor "http" "/.rpc/ListApps" model.serverAddr)
        (list string)
        |> Http.send
            (\resp ->
                case resp of
                    Ok apps ->
                        AppList apps

                    Err err ->
                        RpcError err
            )


appDomain : String -> ServerAddr -> ServerAddr
appDomain subdomain addr =
    { addr | rootHost = subdomain ++ "." ++ addr.rootHost }


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        AppList apps ->
            ( { model | apps = apps }
            , Cmd.none
            )

        RpcError err ->
            ( { model | apps = [ "Rpc Error " ++ toString err ] }
            , Cmd.none
            )

        NewData str ->
            Debug.log str ( model, Cmd.none )


view : Model -> Html Msg
view model =
    let
        viewAppLink appName =
            li [] [ a [ href (urlFor "http" "" (appDomain appName model.serverAddr)) ] [ text appName ] ]
    in
        ul []
            (List.map viewAppLink model.apps)


subscriptions : Model -> Sub Msg
subscriptions model =
    WebSocket.listen (urlFor "ws" "rpc" model.serverAddr) NewData


init : ( Model, Cmd Msg )
init =
    let
        model =
            { apps = []
            , serverAddr =
                -- TODO: extract these from the page:
                { rootHost = "local.sandstorm.io:8000"
                , secure = False
                }
            }
    in
        ( model
        , sendRpc model ListApps
        )


main =
    Html.program
        { init = init
        , view = view
        , update = update
        , subscriptions = subscriptions
        }
