package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/maintainerd/core/internal/dto"
	"github.com/maintainerd/core/internal/model"
	"github.com/maintainerd/core/internal/service"
	"github.com/maintainerd/core/internal/util"
)

type OrganizationHandler struct {
	service service.OrganizationService
}

func NewOrganizationHandler(service service.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{service}
}

func (h *OrganizationHandler) Create(c *gin.Context) {
	var organization model.Organization
	if err := c.ShouldBindJSON(&organization); err != nil {
		util.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	if err := h.service.Create(&organization); err != nil {
		util.Error(c, http.StatusInternalServerError, "Failed to create organization", err.Error())
		return
	}

	util.Created(c, dto.ToOrganizationDTO(&organization), "Organization created successfully")
}

func (h *OrganizationHandler) GetAll(c *gin.Context) {
	organizations, err := h.service.GetAll()
	if err != nil {
		util.Error(c, http.StatusInternalServerError, "Failed to fetch organizations", err.Error())
		return
	}

	var organizationDTOs []dto.OrganizationDTO
	for _, organization := range organizations {
		organizationDTOs = append(organizationDTOs, dto.ToOrganizationDTO(&organization))
	}

	util.Success(c, organizationDTOs, "Organizations fetched successfully")
}

func (h *OrganizationHandler) GetByUUID(c *gin.Context) {
	organizationUUID, err := uuid.Parse(c.Param("organization_uuid"))
	if err != nil {
		util.Error(c, http.StatusBadRequest, "Invalid organization UUID")
		return
	}

	organization, err := h.service.GetByUUID(organizationUUID)
	if err != nil {
		util.Error(c, http.StatusNotFound, "Organization not found")
		return
	}

	util.Success(c, dto.ToOrganizationDTO(organization), "Organization fetched successfully")
}

func (h *OrganizationHandler) Update(c *gin.Context) {
	organizationUUID, err := uuid.Parse(c.Param("organization_uuid"))
	if err != nil {
		util.Error(c, http.StatusBadRequest, "Invalid organization UUID")
		return
	}

	var organization model.Organization
	if err := c.ShouldBindJSON(&organization); err != nil {
		util.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	if err := h.service.UpdateByUUID(organizationUUID, &organization); err != nil {
		util.Error(c, http.StatusInternalServerError, "Failed to update organization", err.Error())
		return
	}

	organization.OrganizationUUID = organizationUUID
	util.Success(c, dto.ToOrganizationDTO(&organization), "Organization updated successfully")
}

func (h *OrganizationHandler) Delete(c *gin.Context) {
	organizationUUID, err := uuid.Parse(c.Param("organization_uuid"))
	if err != nil {
		util.Error(c, http.StatusBadRequest, "Invalid organization UUID")
		return
	}

	if err := h.service.DeleteByUUID(organizationUUID); err != nil {
		util.Error(c, http.StatusInternalServerError, "Failed to delete organization", err.Error())
		return
	}

	util.Success(c, nil, "Organization deleted successfully")
}
