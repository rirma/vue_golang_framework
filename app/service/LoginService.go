package service

import (
    "fmt"
    "time"
    "strconv"
    "app/model"
    "golang.org/x/crypto/bcrypt"
    "github.com/gin-gonic/gin"
    "app/env"
)

type LoginService struct {}

//ログイン処理 戻り値の1つ目は処理が正常に終了したかどうか、2つ目はエラー詳細を返す
func (LoginService) Login(c *gin.Context, email string, password string) (bool, error) {
    //ユーザ情報が一致するか判断
    user := model.User{}

    hash, hash_err := bcrypt.GenerateFromPassword([]byte(password), 10)
    if hash_err != nil {
        fmt.Println(hash_err)
        return false, hash_err
    }
    fmt.Println(string(hash))
    result, db_err := DbEngine.Where("email = ?", email).Get(&user)
    if db_err != nil {
        fmt.Println(db_err)
        return false, db_err
    }
	if !result {
        fmt.Println("Not Found")
        return false, db_err
	}

    //パスワードが正しいか検証
    password_err := bcrypt.CompareHashAndPassword(hash, []byte(password))
    if password_err != nil {
        fmt.Println("パスワードが正しくありません")
        return false, password_err
    }

    //セッションを生成して返す
    login_session, str_err := MakeRandomStr(64)
    if str_err != nil {
        fmt.Println(str_err)
        return false, str_err
    }
    db_session := model.Session{}
    db_session.UserId = user.Id
    db_session.Payload = login_session
    db_session.LastActivity = time.Now().Unix()

    _, session_err := DbEngine.Insert(db_session)
    if session_err != nil {
        return false, session_err
    }

    session_lifetime, _ := strconv.Atoi(env.Get("SESSION_LIFETIME"))
    c.SetCookie("login_session", login_session, session_lifetime, "/", env.Get("APP_URL"), true, true)

    return true, nil
}

//ログアウト処理 戻り値の1つ目は処理が正常に終了したかどうか、2つ目はエラー詳細を返す
func (LoginService) Logout(c *gin.Context) bool {
    //セッション取得
    login_session, user_cookie_err := c.Cookie("login_session")
    if user_cookie_err != nil {
        fmt.Println(user_cookie_err)
        return false
    }

    //セッション検証
    session := model.Session{}
    session_result, db_err := DbEngine.Where("payload = ?", login_session).Get(&session)
    if db_err != nil {
        fmt.Println(db_err)
        panic("session get error")
    }
    if !session_result {
        fmt.Println("Session User Not Found")
        return false
	}
    DbEngine.Delete(&session)

    c.SetCookie("login_session", "", -1, "/", env.Get("APP_URL"), true, true)

    return true
}


//ユーザモデルと取得できたかどうかを返す
func GetLoginUserInformation(c *gin.Context) (model.User, bool) {
    //セッション取得
    login_session, user_cookie_err := c.Cookie("login_session")
    if user_cookie_err != nil {
        fmt.Println(user_cookie_err)
        return model.User{}, false
    }

    //セッション検証
    session := model.Session{}
    session_result, db_err := DbEngine.Where("payload = ?", login_session).Get(&session)
    if db_err != nil {
        fmt.Println(db_err)
        panic("session get error")
    }
    if !session_result {
        fmt.Println("Session User Not Found")
        return model.User{}, false
	}

    //セッションに該当するユーザ返却
    user := model.User{}

    user_result, user_db_err := DbEngine.Where("id = ?", session.UserId).Get(&user)
    if user_db_err != nil {
        fmt.Println(user_db_err)
        panic("user get error")
    }
	if !user_result {
        fmt.Println("User Not Found")
        return model.User{}, false
	}

    return user, true
}