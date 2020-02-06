package tests

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/deliriumproducts/aumo"
	"github.com/deliriumproducts/aumo/mysql"
	"github.com/deliriumproducts/aumo/receipt"
	"github.com/stretchr/testify/require"
)

func TestReceiptService(t *testing.T) {
	sess, err := SetupDB()
	if err != nil {
		t.Error(err)
	}

	// Cleanup
	defer func() {
		TidyDB(sess)
		sess.Close()
	}()

	ustore := mysql.NewUserStore(sess)
	rstore := mysql.NewReceiptStore(sess)

	rs := receipt.New(rstore, ustore)

	t.Run("create_receipt", func(t *testing.T) {
		defer TidyDB(sess)

		receipt := aumo.NewReceipt("Paconi: 230")

		err = rs.Create(receipt)
		require.Nil(t, err, "shouldn't return an error")

		gotReceipt, err := rstore.FindByID(nil, receipt.ReceiptID.String())
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, *receipt, *gotReceipt)
	})

	t.Run("get_receipt", func(t *testing.T) {
		defer TidyDB(sess)

		receipt := createReceipt(t, rstore)

		gotReceipt, err := rs.Receipt(receipt.ReceiptID.String())

		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, *receipt, *gotReceipt)
	})

	t.Run("get_receipts", func(t *testing.T) {
		defer TidyDB(sess)

		receipts := []aumo.Receipt{
			*aumo.NewReceipt(faker.AmountWithCurrency()),
			*aumo.NewReceipt(faker.AmountWithCurrency()),
			*aumo.NewReceipt(faker.AmountWithCurrency()),
		}

		for _, receipt := range receipts {
			err := rstore.Save(nil, &receipt)
			require.Nil(t, err, "shouldn't return an error")
		}

		gotReceipts, err := rs.Receipts()
		require.Nil(t, err, "shouldn't return an error")
		require.ElementsMatch(t, gotReceipts, receipts, "should be equal")
	})

	t.Run("delete_receipt", func(t *testing.T) {
		defer TidyDB(sess)

		receipt := createReceipt(t, rstore)

		err = rs.Delete(receipt.ReceiptID.String())
		require.Nil(t, err, "shouldn't return an error")

		_, err = rstore.FindByID(nil, receipt.ReceiptID.String())
		require.Equal(t, err, aumo.ErrReceiptNotFound)
	})

	t.Run("update_receipt", func(t *testing.T) {
		defer TidyDB(sess)

		receipt := createReceipt(t, rstore)
		receipt.Content = "Kaufland 23233232323"

		err = rs.Update(receipt.ReceiptID.String(), receipt)
		require.Nil(t, err, "shouldn't return an error")

		gotReceipt, err := rstore.FindByID(nil, receipt.ReceiptID.String())
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, *receipt, *gotReceipt)
	})

	t.Run("claim_receipt", func(t *testing.T) {
		defer TidyDB(sess)

		user := createUser(t, ustore)

		t.Run("valid", func(t *testing.T) {
			receipt := createReceipt(t, rstore)
			require.Equal(t, false, receipt.IsClaimed())

			var err error
			receipt, err = rs.ClaimReceipt(user.ID.String(), receipt.ReceiptID.String())
			require.Nil(t, err, "shouldn't return an error")
			require.Equal(t, true, receipt.IsClaimed())

			user.Points += aumo.UserPointsPerReceipt

			gotReceipt, err := rstore.FindByID(nil, receipt.ReceiptID.String())
			require.Nil(t, err, "shouldn't return an error")
			require.Equal(t, true, gotReceipt.IsClaimed())

			gotUser, err := ustore.FindByID(nil, user.ID.String(), true)
			require.Nil(t, err, "shouldn't return an error")
			require.Contains(t, gotUser.Receipts, *gotReceipt)
			require.Equal(t, user.Points, gotUser.Points)
		})
	})
}
