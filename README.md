# milter-sample

メールをトレースするためのMilterです。Postfixのみで試してます。

## 使い方

```
# Milterの設定
smtpd_milters = inet:localhost:12345
non_smtpd_milters = inet:localhost:12345
milter_protocol = 6
milter_default_action = accept
```
