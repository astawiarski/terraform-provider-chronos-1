package chronos

import (
	"fmt"
	"github.com/behance/go-chronos/chronos"
	//"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"log"
	//"reflect"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func readExampleAppConfiguration() string {
	time.Sleep(1 * time.Second)
	bytes, err := ioutil.ReadFile("../test/example.tf")
	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}

func TestAccChronosApp_basic(t *testing.T) {

	var j chronos.Jobs

	testCheckCreate := func(jobs *chronos.Jobs) resource.TestCheckFunc {
		return func(s *terraform.State) error {
			time.Sleep(1 * time.Second)
			for _, j := range *jobs {
				if j.Name == "" {
					return fmt.Errorf("Didn't return a Name so something is broken: %#v", j)
				}
			}
			return nil
		}
	}

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: readExampleAppConfiguration(),
				Check: resource.ComposeTestCheckFunc(
					testAccReadApp("chronos_job.job-create-example", &j),
					testCheckCreate(&j),
				),
			},
		},
	})
}

func testAccReadApp(name string, jobs *chronos.Jobs) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		time.Sleep(1 * time.Second)
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("chronos_job resource not found: %s", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("chronos_job id not set correctly: %s", name)
		}

		log.Printf("=== testAccContainerExists: rs ===\n%#v\n", rs)

		config := testAccProvider.Meta().(config)
		client := config.Client

		time.Sleep(1 * time.Second)
		jobRead, err := client.Jobs()

		log.Printf("=== testAccContainerExists: appRead ===\n%#v\n", jobRead)

		*jobs = *jobRead

		return err
	}
}

func testAccCheckChronosAppDestroy(s *terraform.State) error {
	time.Sleep(1 * time.Second)
	config := testAccProvider.Meta().(config)
	client := config.Client

	jobs, err := client.Jobs()
	if len(*jobs) > 0 {
		return fmt.Errorf("Job not deleted! %#v", err)
	}
	return err
}

//
//func TestEndChronosConvertion(t *testing.T) {
//
//	client, err := MockClient(chronos.Jobs{})
//
//	if err != nil {
//		t.Error("Couldn't get client %s", err)
//	}
//	jobs, errj := client.Jobs()
//	if errj != nil {
//		t.Error("Couldn't get Jobs %s", err)
//	}
//	t.Log(len(*jobs))
//	for _, j := range *jobs {
//		var d, d2 schema.ResourceData
//		t.Log(d)
//		d.Set("name", "fail")
//		jobToResource(j, &d) // remove all autosetted value
//		if v, ok := d.GetOk("name"); ok {
//			t.Log(v.(string))
//		}
//		jobToResource(*resourceToJob(&d), &d2)
//		t.Log(j.Name)
//		t.Log(resourceToJob(&d).Name)
//		if j.Name != resourceToJob(&d).Name || !reflect.DeepEqual(d, d2) {
//			t.Fatal(d, "not the same as ", d2)
//		}
//	}
//
//}
