package provider

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// defaultClusterTimeout is the amount of time that can elapse
// before a cluster read operation should timeout.
var defaultVaultClusterTimeout = time.Minute * 5

// createTimeout is the amount of time that can elapse
// before a cluster create operation should timeout.
var createVaultClusterTimeout = time.Minute * 35

// deleteTimeout is the amount of time that can elapse
// before a cluster delete operation should timeout.
var deleteVaultClusterTimeout = time.Minute * 25

func resourceVaultCluster() *schema.Resource {
	return &schema.Resource{
		Description:   "The Vault cluster resource allows you to manage an HCP Vault cluster.",
		CreateContext: resourceVaultClusterCreate,
		ReadContext:   resourceVaultClusterRead,
		DeleteContext: resourceVaultClusterDelete,
		Timeouts: &schema.ResourceTimeout{
			Create:  &createVaultClusterTimeout,
			Delete:  &deleteVaultClusterTimeout,
			Default: &defaultVaultClusterTimeout,
		},
		Schema: map[string]*schema.Schema{
			// Required inputs
			"cluster_id": {
				Description:      "The ID of the HCP Vault cluster.",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validateSlugID,
			},
			"hvn_id": {
				Description:      "The ID of the HVN this HCP Vault cluster is associated to.",
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validateSlugID,
			},
			// optional fields
			"public_endpoint": {
				Description: "Denotes that the cluster has a public endpoint for the Vault UI. Defaults to false.",
				Type:        schema.TypeBool,
				Default:     false,
				Optional:    true,
				ForceNew:    true,
			},
			"min_vault_version": {
				Description:      "The minimum Vault version to use when creating the cluster. If not specified, it is defaulted to the version that is currently recommended by HCP.",
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validateSemVer,
				ForceNew:         true,
			},
			// computed outputs
			// TODO: once more tiers are supported and can be changed by users, make this a required input.
			"tier": {
				Description: "The tier that the HCP Vault cluster will be provisioned as.  Only 'development' is available at this time.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"organization_id": {
				Description: "The ID of the organization this HCP Vault cluster is located in.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"project_id": {
				Description: "The ID of the project this HCP Vault cluster is located in.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
			},
			"cloud_provider": {
				Description: "The provider where the HCP Vault cluster is located.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"region": {
				Description: "The region where the HCP Vault cluster is located.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"namespace": {
				Description: "The name of the customer namespace this HCP Vault cluster is located in.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"vault_version": {
				Description: "The Vault version of the cluster.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"vault_public_endpoint_url": {
				Description: "The public URL for the Vault UI. This will be empty if `public_endpoint` is `false`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"vault_private_endpoint_url": {
				Description: "The private URL for the Vault UI.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"created_at": {
				Description: "The time that the Vault cluster was created.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceVaultClusterCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceVaultClusterRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceVaultClusterDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
