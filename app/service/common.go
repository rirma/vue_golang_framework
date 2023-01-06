package service

import (
    "errors"
    "math/rand"
)

func MakeRandomStr(digit uint32) (string, error) {
    // 乱数を生成
    b := make([]byte, digit)
    if _, err := rand.Read(b); err != nil {
        return "", errors.New("unexpected error...")
    }

    // 乱数をASCII文字列に変換
    var result string
    for _, v := range b {
        // 制御文字が当たらないように調整
        result += string(v%byte(94) + 33)
    }
    return result, nil
}