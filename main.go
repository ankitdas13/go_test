package main

import (
	"fmt"
	"encoding/json"
	razorpay "github.com/razorpay/razorpay-go"
)

func main(){
	client := razorpay.NewClient("rzp_test_k6uL897VPBz20q", "EnLs21M47BllR3X8PSFtjtbd")

	data := map[string]interface{}{
		"amount": 50000,
		"currency": "INR",
		"receipt": "some_receipt_id",
		"partial_payment": false,
		"notes": map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		  },
	  }
	body, err := client.Order.Create(data, nil)
	  

	jsonString, err1 := json.Marshal(body)
    fmt.Println(string(jsonString))
    fmt.Println(err1)
	fmt.Println(err)
}

//Direct transfer (sdk audit) : This feature is not enabled for this merchant.



// transfer
//rzp_live_6nctVu1OSjHGSE  rDOPqzxgxatD3pGm7HOSA3Ey

// Fetch transfers for a settlement

// ```go
// data:= map[string]interface{}{
//   "recipient_settlement_id": "<recipientSettlementId>",
// }
// body, err := client.Transfer.All(data, nil)


//1. Customer already exists for the merchant
//2. {"amount":0,"amount_due":0,"amount_paid":0,"attempts":0,"created_at":1656873250,"currency":"INR","entity":"order","id":"order_Jozfws8G585Jhj","notes":{"notes_key_1":"Beam me up Scotty","notes_key_2":"Engage"},"offer_id":null,"offers":{"count":0,"entity":"collection","items":[]},"payments":{"count":0,"entity":"collection","items":[]},"receipt":"Receipt No. 1","status":"created","token":{"auth_type":"physical","bank_account":{"account_number":"11214311215411","account_type":"savings","bank_name":"HDFC Bank","beneficiary_email":"gorav.kumar@example.com","beneficiary_mobile":"9123456780","ifsc":"HDFC0000001","name":"Gaurav Kumar"},"currency":"INR","expire_at":2709971120,"failure_reason":null,"first_payment_amount":10000,"max_amount":10000000,"method":"nach","nach":{"create_form":true,"description":"Paper NACH Gaurav Kumar","form_reference1":"Recurring Payment for Gaurav Kumar","form_reference2":"Method Paper NACH","prefilled_form":"https://rzp.io/i/KtpxTztI5w","prefilled_form_transient":"https://rzp.io/i/xKJv3LU","upload_form_url":"https://rzp.io/i/UL8UJO2Yk"},"notes":{"notes_key_1":"Tea, Earl Grey, Hot","notes_key_2":"Tea, Earl Grey… decaf."},"recurring_status":null}}
//3. receipt must be unique for each item : Receipt No. 1, expire_by should be at least 15 minutes after current time, expire_by should be at least 15 minutes after current time
//4. 