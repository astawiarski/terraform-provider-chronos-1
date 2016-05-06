package chronos

import (
	"errors"
	"fmt"
	"github.com/behance/go-chronos/chronos"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"time"
)

func resourceChronosJob() *schema.Resource {
	return &schema.Resource{
		Create: resourceChronosJobCreate,
		Read:   resourceChronosJobRead,
		Update: resourceChronosJobUpdate,
		Delete: resourceChronosJobDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"command": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"shell": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: false,
			},
			"epsilon": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"executor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"executor_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
				ForceNew: false,
			},
			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"owner_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"async": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: false,
			},
			"success_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
				ForceNew: false,
			},
			"error_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
				ForceNew: false,
			},
			"last_success": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"last_error": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"cpus": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				ForceNew: false,
				Default:  1,
			},
			"disk": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				ForceNew: false,
				Default:  64,
			},
			"mem": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				ForceNew: false,
				Default:  64,
			},
			"disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: false,
			},
			"soft_error": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: false,
			},
			"data_processing_job_type": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: false,
			},
			"error_since_last_success": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
				ForceNew: false,
			},
			"uris": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"env": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: false,
			},
			"arguments": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"high_priority": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: false,
			},
			"run_as_user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"container": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"image": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"network": &schema.Schema{
							Type:     schema.TypeString,
							Default:  "HOST",
							Optional: true,
						},
						//"volumes": &schema.Schema{
						//	Type:     schema.TypeMap,
						//	Optional: true,
						//	ForceNew: false,
						//},
						"force_pull": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
							ForceNew: false,
						},
						"parameters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: false,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"parameter": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: false,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": &schema.Schema{
													Type:     schema.TypeString,
													Default:  "tcp",
													Optional: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Default:  "tcp",
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"schedule_timezone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			// TODO constraints
			"parents": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceChronosJobCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(config)
	client := config.Client

	job := resourceToJob(d)
	var err error
	if len(job.Parents) > 0 {
		err = client.AddDependentJob(job)
	} else {
		err = client.AddScheduledJob(job)
	}
	if err != nil {
		log.Println("[ERROR] creating jobs", err)
		return err
	}
	d.SetId(job.Name)
	d.Partial(false)

	return resourceChronosJobRead(d, meta)
}

func resourceChronosJobRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(config)
	client := config.Client
	var jobs *chronos.Jobs
	time.Sleep(time.Second)
	jobs, err := client.Jobs()

	//log.Print("############", len(*jobs))
	if err != nil {
		log.Println("[ERROR] read jobs", err)
		return err
	}

	for _, job := range *jobs {
		log.Print("############", d.Id(), job.Name)
		if d.Id() == job.Name {
			return jobToResource(job, d)
		}
	}

	return errors.New("Jobs not found")
}

func resourceChronosJobUpdate(d *schema.ResourceData, meta interface{}) error {

	d.Partial(true)
	err := resourceChronosJobDelete(d, meta)
	if err != nil {
		log.Println("[Warning] Deleted Jobs failed", err)
	}

	err = resourceChronosJobCreate(d, meta)
	if err != nil {
		d.Partial(false)
	}

	return err
}

func resourceChronosJobDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(config)
	client := config.Client
	err := client.DeleteJob(d.Id())
	if err != nil {
		d.Partial(true)
	}
	return err
}

func resourceToJob(d *schema.ResourceData) *chronos.Job {

	job := new(chronos.Job)

	if v, ok := d.GetOk("name"); ok {
		job.Name = v.(string)
	}
	if v, ok := d.GetOk("command"); ok {
		job.Command = v.(string)
	}

	if v, ok := d.GetOk("shell"); ok {
		job.Shell = v.(bool)
	}

	if v, ok := d.GetOk("epsilon"); ok {
		job.Epsilon = v.(string)
	}

	if v, ok := d.GetOk("executor"); ok {
		job.Executor = v.(string)
	}

	if v, ok := d.GetOk("executor_flags"); ok {
		job.ExecutorFlags = v.(string)
	}

	if v, ok := d.GetOk("retries"); ok {
		job.Retries = v.(int)
	}
	if v, ok := d.GetOk("owner"); ok {
		job.Owner = v.(string)
	}

	if v, ok := d.GetOk("owner_name"); ok {
		job.OwnerName = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		job.Description = v.(string)
	}

	if v, ok := d.GetOk("async"); ok {
		job.Async = v.(bool)
	}

	// sucess count
	// error count
	// last success
	// last error

	if v, ok := d.GetOk("cpus"); ok {
		job.CPUs = float32(v.(float64))
	}

	if v, ok := d.GetOk("disk"); ok {
		job.Disk = float32(v.(float64))
	}

	if v, ok := d.GetOk("mem"); ok {
		job.Mem = float32(v.(float64))
	}

	if v, ok := d.GetOk("disable"); ok {
		job.Disabled = v.(bool)
	}

	if v, ok := d.GetOk("soft_error"); ok {
		job.SoftError = v.(bool)
	}

	if v, ok := d.GetOk("data_processing_job_type"); ok {
		job.DataProcessingJobType = v.(bool)
	}

	// ErrorsSinceLastSuccess

	if v, ok := d.GetOk("uris.#"); ok {
		uris := make([]string, v.(int))

		for i := range uris {
			uris[i] = d.Get("uris." + strconv.Itoa(i)).(string)
		}

		if len(uris) != 0 {
			job.URIs = uris
		}
	}

	if v, ok := d.GetOk("env"); ok {
		maptmp := v.(map[string]interface{})
		var env []map[string]string
		for name, value := range maptmp {
			if str, okStr := value.(string); okStr {
				env = append(env, map[string]string{
					"value": str,
					"name":  name,
				})
			}
		}

		job.EnvironmentVariables = env
	} else {
		job.EnvironmentVariables = make([]map[string]string, 0)
	}

	if v, ok := d.GetOk("arguments.#"); ok {
		arguments := make([]string, v.(int))

		for i := range arguments {
			arguments[i] = d.Get("arguments." + strconv.Itoa(i)).(string)
		}

		if len(arguments) != 0 {
			job.Arguments = arguments
		}
	}

	if v, ok := d.GetOk("high_priority"); ok {
		job.HighPriority = v.(bool)
	}

	if v, ok := d.GetOk("run_as_user"); ok {
		job.RunAsUser = v.(string)
	}

	if _, ok := d.GetOk("container"); ok {
		var containerType, containerNetwork, containerImage string
		if v, ok := d.GetOk("container.0.type"); ok {
			containerType = v.(string)
		}

		if v, ok := d.GetOk("container.0.network"); ok {
			containerNetwork = v.(string)
		}

		if v, ok := d.GetOk("container.0.image"); ok {
			containerImage = v.(string)
		}

		//var container_volumes []map[string]string
		//if v, ok := d.GetOk("container.0.volumes.#"); ok {
		//	container_volumes = make([]map[string]string, v.(int))
		//	for i, _ := range container_volumes {
		//		container_volumes[i] = d.Get(fmt.Sprintf("container.0.volumes.%d", i)).(map[string]string)
		//	}
		//} else {
		//	container_volumes = make([]map[string]string, 0)
		//}

		var forcePull bool
		if v, ok := d.GetOk("container.0.force_pull"); ok {
			forcePull = v.(bool)
		}

		var containerParameters []map[string]string
		if v, ok := d.GetOk("container.0.parameters.0.parameter.#"); ok {
			containerParameters = make([]map[string]string, v.(int))
			for i := range containerParameters {
				paramMap := d.Get(
					fmt.Sprintf("container.0.parameters.0.parameter.%d", i),
				).(map[string]interface{})
				containerParameters[i] = map[string]string{
					"key":   paramMap["key"].(string),
					"value": paramMap["value"].(string),
				}
			}
		}

		job.Container = &chronos.Container{
			Type:    containerType,
			Network: containerNetwork,
			Image:   containerImage,
			//Volumes:        container_volumes,
			ForcePullImage: forcePull,
			Parameters:     containerParameters,
		}
	}

	if v, ok := d.GetOk("schedule"); ok {
		job.Schedule = v.(string)
	}

	if v, ok := d.GetOk("schedule_timezone"); ok {
		job.ScheduleTimeZone = v.(string)
	}

	// TODO constraints

	if v, ok := d.GetOk("parents.#"); ok {
		parents := make([]string, v.(int))

		for i := range parents {
			parents[i] = d.Get("parents." + strconv.Itoa(i)).(string)
		}

		if len(parents) != 0 {
			job.Parents = parents
		}
	}

	return job
}

func jobToResource(job chronos.Job, d *schema.ResourceData) error {

	d.SetId(job.Name)
	d.Set("name", job.Name)
	d.SetPartial("name")

	d.Set("command", job.Command)
	d.SetPartial("command")

	d.Set("shell", job.Shell)
	d.SetPartial("shell")

	d.Set("epsilon", job.Epsilon)
	d.SetPartial("epsilon")

	d.Set("executor", job.Executor)
	d.SetPartial("executor")

	d.Set("executor_flags", job.ExecutorFlags)
	d.SetPartial("executor_flags")

	d.Set("retries", job.Retries)
	d.SetPartial("retries")

	d.Set("owner", job.Owner)
	d.SetPartial("owner")

	d.Set("owner_name", job.OwnerName)
	d.SetPartial("owner_name")

	d.Set("description", job.Description)
	d.SetPartial("description")

	d.Set("async", job.Async)
	d.SetPartial("async")

	// sucess count
	// error count
	// last success
	// last error

	d.Set("cPUs", job.CPUs)
	d.SetPartial("cPUs")

	d.Set("disk", job.Disk)
	d.SetPartial("disk")

	d.Set("mem", job.Mem)
	d.SetPartial("mem")

	d.Set("disabled", job.Disabled)
	d.SetPartial("disabled")

	d.Set("soft_error", job.SoftError)
	d.SetPartial("soft_error")

	d.Set("data_processing_job_type", job.DataProcessingJobType)
	d.SetPartial("data_processing_job_type")

	// ErrorsSinceLastSuccess

	d.Set("uris", job.URIs)
	d.SetPartial("uris")

	//       []map[string]string
	env := make(map[string]string)
	for _, e := range job.EnvironmentVariables {
		env[e["name"]] = e["value"]
	}
	d.Set("env", env)
	d.SetPartial("env")

	d.Set("arguments", job.Arguments)
	d.SetPartial("arguments")

	d.Set("high_priority", job.HighPriority)
	d.SetPartial("high_priority")

	d.Set("run_as_user", job.RunAsUser)
	d.SetPartial("run_as_user")
	//                  *Container
	if job.Container != nil {
		d.Set("container", map[string]interface{}{
			"type":    job.Container.Type,
			"image":   job.Container.Image,
			"network": job.Container.Network,
			//		"volumes":    job.Container.Volumes,
			"force_pull": job.Container.ForcePullImage,
			"parameters": job.Container.Parameters,
		})
		d.SetPartial("container")
	}
	d.Set("schedule", job.Schedule)
	d.SetPartial("schedule")

	d.Set("schedule_timezone", job.ScheduleTimeZone)
	d.SetPartial("schedule_timezone")

	//  TODO Constraints               []map[string]string

	d.Set("parents", job.Parents)
	d.SetPartial("parents")

	return nil
}
