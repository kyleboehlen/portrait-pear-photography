package auth

// Bcrypt hashes for admin password, rotated monthly because I can.
// Format: YYYY-MM
// 420 bits of entropy + bcrypt cost of 17. You do the math, I'll take those odds.
var theForbiddenHashArray = map[string]string{
	"2025-01": "$2a$17$0iMVFePIFc0jLFzKTkdVEuWCVJ6lR/FrmyjmyNT5R1Ut20CZmFsBi",
	"2025-02": "$2a$17$8TP4eBs9lxkgyJ.HISeyH.38VkrMHh.unIiPFtpk8FZ37IyN4F.s.",
	"2025-03": "$2a$17$/8RmzfZhGWbfQglW/yYTp.MajNIFrgaXu8TEMCAFOMbOtfA2W2Cv2",
	"2025-04": "$2a$17$hGYJWFU4HqbtvU5EUsP8d.5D1yQBJZ228Pk/z6.qEv3b/lcxIzxb2",
	"2025-05": "$2a$17$/OjaeFHQJ7k1yD72A7V2X.8hz5MpcyiXIWMCdrJlFciv1q8EUmVl2",
	"2025-06": "$2a$17$2uflpk8QdHQ4hl7Z1n8N7.h5wkrCO79Inz6SW4YDs/XcKW7nATnzu",
	"2025-07": "$2a$17$m6ymoHv4fSkLOWlrPZjTzemJs6rEItjQZ8sr9gDq6GPUZrid0LWBy",
	"2025-08": "$2a$17$YD8wCvuBU7YbeR7Q2juQve8FiaAJ5L35YGNlQr.aY7PcnVriA1EPS",
	"2025-09": "$2a$17$V7I2MZg2nnaVVg9e3qts1.PwSU22fwnJ1YTPpWE51565X7SY.Bj8m",
	"2025-10": "$2a$17$UuBp/lvjWQfqNsHdpoIN4O9y/obMIJkbfndvNxnJ7JfXa3Kso/UOa",
	"2025-11": "$2a$17$RXuLKyzmL9b/LOOV5WvQJeDZSa8t9rqGQS.QGh6ZcDCHWZVUEuX7S",
	"2025-12": "$2a$17$psRfojEKZq8XtOPEaHgx0eM5RL9LfKl15E5dKeJcqO0xgbG6nWVO.",
	"2026-01": "$2a$17$x4TRTSwBtAFljlOfAWFg1ek7qL1zj9EVOUff56MJg2ZigsA1h6cdC",
	"2026-02": "$2a$17$XnzCr92MpdVmrUr6Z/0OD.r32FWagF9xNDZgV98jxLP547xGh4aoi",
	"2026-03": "$2a$17$VvpGgX1GgHZuDOVmaWysPuMmdtm2Wfy7ux1LOVjwF81rN7Px.Z6Fy",
	"2026-04": "$2a$17$rJ1E/kw/ANfp6G8lKDReYO50u9DbG3116rdgFu3MwFyKHc/mCJLUS",
	"2026-05": "$2a$17$gsFLRmH.X9/G2HYO0xzUfeIGg4nGbh71Ykmn9w34cZ90C6jjaFoK.",
	"2026-06": "$2a$17$NFQ1XfGmS2Kj595PI4f2FurOLhZOOKVhsppa1KtrXU3pJFETeJsQC",
	"2026-07": "$2a$17$T5R4txDROsa4NS3ensD6ieWxa479HxWsagEZBY4xQJQgTvKGjq112",
	"2026-08": "$2a$17$NVV1.3BQHyTXWr5Yd6b5FuvPRosyA4.Ls8JJ6X6OdfLKKogllyEve",
}
