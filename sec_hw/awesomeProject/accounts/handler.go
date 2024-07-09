package accounts

import (
	"awesomeProject/accounts/dto"
	"awesomeProject/accounts/models"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

func New() *Handler {
	return &Handler{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

// Create account with name 'name' and balance 'amount'
func (h *Handler) CreateAccount(c echo.Context) error {
	var request dto.CreateAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.Lock()
	defer h.guard.Unlock()

	if _, ok := h.accounts[request.Name]; ok {
		return c.String(http.StatusForbidden, "account already exists")
	}

	h.accounts[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: request.Amount,
	}

	return c.NoContent(http.StatusOK)
}

// Get account with name 'name'
func (h *Handler) GetAccount(c echo.Context) error {
	name := c.QueryParams().Get("name")

	h.guard.RLock()
	account, ok := h.accounts[name]
	h.guard.RUnlock()

	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	response := dto.GetAccountResponse{
		Name:   account.Name,
		Amount: account.Amount,
	}

	return c.JSON(http.StatusOK, response)
}

// Delete account with name 'name'
func (h *Handler) DeleteAccount(c echo.Context) error {
	var request dto.DeleteAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	h.guard.Lock()
	defer h.guard.Unlock()

	if _, ok := h.accounts[request.Name]; !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	delete(h.accounts, request.Name)

	return c.NoContent(http.StatusOK)
}

// Set balance 'amount' on account with name 'name'
func (h *Handler) SetBalance(c echo.Context) error {
	var request dto.SetBalanceRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	h.guard.Lock()
	defer h.guard.Unlock()

	if _, ok := h.accounts[request.Name]; !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	h.accounts[request.Name].Amount = request.Amount

	return c.NoContent(http.StatusOK)
}

// Change name to 'new_name' on account with name 'name'
func (h *Handler) RenameAccount(c echo.Context) error {
	var request dto.RenameAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	h.guard.Lock()
	defer h.guard.Unlock()

	if _, ok := h.accounts[request.Name]; !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	if _, ok := h.accounts[request.NewName]; ok {
		return c.String(http.StatusNotFound, "account with new name already exists")
	}

	h.accounts[request.NewName] = h.accounts[request.Name]
	delete(h.accounts, request.Name)

	return c.NoContent(http.StatusOK)
}
