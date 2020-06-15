package hashid

import hashids "github.com/speps/go-hashids"

/*
https://mp.weixin.qq.com/s/b3a2_jIeGK0QK5ikqiz8yw

## 问题

为什么要对 int 类型的数据加密，它的应用场景是什么？

比如：有一个商品详情界面 URL 为 /product/1001，这种情况很容易被别人猜测，
比如输入 /product/1002、/product/1003 尝试着去查看详情，这样的话信息就暴露了，
如果别人想抓数据的话，只需要将后面的 ID 递增抓取就可以了，怎么解决这个问题？

比如：有一个用户邀请码需求，用户可以将自己的邀请码分享出去，当新用户使用这个邀请码注册的时候，
就会给邀请者和被邀请者双方发奖励，通过 URL /user/1001 注册的，
表示用户ID为 1001 的邀请的，这样用户ID很容易被修改，怎么解决这个问题？

## 分析

上面的两个场景都是需要对 int 类型的数据进行加密，避免 ID 泄露。

需要满足以下特性：

支持自定义 salt，保证加密后的是独一无二。
支持加密和解密。
支持多语言。

*/

/*
https://hashids.org/

Hashids is a small open-source library that generates short, unique, non-sequential ids from numbers.

It converts numbers like 347 into strings like “yr8”, or array of numbers like [27, 986] into “3kTMd”.

You can also decode those ids back.
This is useful in bundling several parameters into one or simply using them as short UIDs.
*/

type HashID struct {
	*hashids.HashID
}

// NewHashID returns a new HashID.
func NewHashID(salt string, minLength int) HashID {
	h, e := NewHashIDE(salt, minLength)
	if e != nil {
		panic(e)
	}

	return h
}

// NewHashIDE returns a new hash ID or error.
func NewHashIDE(salt string, minLength int) (HashID, error) {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength

	h, err := hashids.NewWithData(hd)

	return HashID{HashID: h}, err
}

// Encrypt 加密.
func (h HashID) Encrypt(params ...int) string {
	if s, err := h.EncryptE(params...); err != nil {
		panic(err)
	} else {
		return s
	}
}

// Decrypt  解密.
func (h HashID) Decrypt(hash string) []int {
	if r, err := h.DecryptE(hash); err != nil {
		panic(err)
	} else {
		return r
	}
}

// EncryptE 加密.
func (h HashID) EncryptE(params ...int) (string, error) {
	return h.Encode(params)
}

// DecryptE  解密.
func (h HashID) DecryptE(hash string) ([]int, error) {
	return h.DecodeWithError(hash)
}
