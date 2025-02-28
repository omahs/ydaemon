package partners

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/helpers"
)

// GetPartner will, for a given address on given chainID, return the meta informations for the partner.
func (y Controller) GetPartner(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	partnerAddressOrName := c.Param("addressOrName")
	partner, ok := Store.PartnersByName[chainID][partnerAddressOrName]
	if !ok {
		partner, ok = Store.PartnersByAddress[chainID][addresses.ToMixedcase(partnerAddressOrName)]
		if !ok {
			c.String(http.StatusBadRequest, "no data available")
			return
		}
	}

	// partner = partner.BalanceOf() //WORK IN PROGRESS

	c.JSON(http.StatusOK, partner)
}
