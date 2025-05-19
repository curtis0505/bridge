package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Accounts struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	UUID               string             `bson:"uuid"`                  // Account Universally Unique ID
	Email              string             `bson:"email"`                 // Email Address
	SID                string             `bson:"sid"`                   // Social id
	DID                string             `bson:"did"`                   // device id
	Ci                 string             `bson:"ci"`                    // korea user id
	ActType            string             `bson:"act_type"`              // Account Type google:GGL, apple:APL, facebook:FCB...
	Name               string             `bson:"name"`                  // User Name
	Status             int                `bson:"status"`                // Account Status erd ref.
	IsLogin            bool               `bson:"is_login"`              // User Login Status
	MnemonicHash       string             `bson:"mnemonic_hash"`         // Mnemonic Hash
	IsOTP              bool               `bson:"is_otp"`                // OTP Agreement
	IsTerms            bool               `bson:"is_terms"`              // Terms Agreement
	KycStatus          int                `bson:"kyc_status"`            // KYC Agreement
	Nation             string             `bson:"nation"`                // Nation
	Birth              string             `bson:"birth"`                 // Birth, DOB(Date of Birth)
	ScOtp              string             `bson:"sc_otp"`                // OTP Secure Code
	LoginAt            time.Time          `bson:"login_at"`              // Last Login Time
	UpdateAt           time.Time          `bson:"update_at"`             // Account Info Update Time
	RegAt              time.Time          `bson:"reg_at"`                // Regist Time
	LeavedAt           time.Time          `bson:"leaved_at"`             // 탈퇴 시간
	IsBusiness         int                `bson:"is_business"`           // 사업자 여부
	KycReason          string             `bson:"kyc_reason"`            // KYC 이유
	KycSite            string             `bson:"kyc_site"`              // KYC 인증 업체 정보("argos", "veriff", "")
	AddressType        AddressType        `bson:"address_type"`          // 계정 타입. (0: user, 1: internal, 2:test)
	AppVersion         string             `bson:"app_version"`           // 앱 버전
	OsInfo             string             `bson:"os_info"`               // 디바이스 정보
	EventReferralCode  string             `bson:"event_referral_code"`   // 이벤트 레퍼럴 코드
	SubmissionID       string             `bson:"submission_id"`         // argos kyc submission_id
	ApplicantID        string             `bson:"applicant_id"`          // argos kyc applicant_id
	RiskLevel          int                `bson:"risk_level"`            // argos aml risk level
	AmlStatus          int                `bson:"aml_status"`            // argos aml status
	ArgosVersion       string             `bson:"argos_version"`         // argos version
	KycFinalAt         time.Time          `bson:"kyc_final_at"`          // kyc 어드민 승인시점
	KycUpdateAt        time.Time          `bson:"kyc_update_at"`         // kyc 상태 변경 시점
	KycFastApproveAble bool               `bson:"kyc_fast_approve_able"` // kyc 자동 승인 가능 여부
	KycFastApproved    int                `bson:"kyc_fast_approved"`     // kyc 자동 승인 여부
	Location           string             `bson:"location"`              // 집주소
	IpAddress          string             `bson:"ip_address"`            // kyc 인증 기준 IP
	JoinAddress        string             `bson:"join_address"`          // 가입 주소
}

type AccountsRepository interface {
	FindAccounts(ctx context.Context, filter bson.M, opts ...*options.FindOptions) ([]*Accounts, error)
	FindOneAccounts(ctx context.Context, filter bson.M, opts ...*options.FindOneOptions) (*Accounts, error)
	CountAccounts(ctx context.Context, filter any, opts ...*options.CountOptions) (int64, error)
	UpdateOneAccounts(ctx context.Context, filter, update any) error
}

func (a *AccountDB) FindAccounts(ctx context.Context, filter bson.M, opts ...*options.FindOptions) ([]*Accounts, error) {
	return a.accounts.Find(ctx, filter, opts...)
}

func (a *AccountDB) FindOneAccounts(ctx context.Context, filter bson.M, opts ...*options.FindOneOptions) (*Accounts, error) {
	return a.accounts.FindOne(ctx, filter, opts...)
}

func (a *AccountDB) CountAccounts(ctx context.Context, filter any, opts ...*options.CountOptions) (int64, error) {
	return a.accounts.CountDocuments(ctx, filter, opts...)
}

func (a *AccountDB) UpdateOneAccounts(ctx context.Context, filter, update any) error {
	return a.accounts.UpdateOne(ctx, filter, update)
}
