package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gitee.com/go-package/carbon"
	"github.com/urfave/cli/v2"
)

func main() {

	app := cli.NewApp()
	app.Name = "lunargo"
	app.Usage = "生成农历及时间相关处理"
	app.Description = "生成农历及时间相关处理"
	// s 表示太阳历 sy：年 sm：月 sd：日 sfull：全量 lw：星期
	// l 表示太阴历 ly：年 lm：月 ld：日 lfull：全量 la：属相
	app.UsageText = "./lunargo [-d day] <-o <sy|sm|sd|sw|ly|lm|ld|la|lfull|sfull>>"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "day",
			Aliases: []string{"d"},
			Usage:   "日期",
		},
		&cli.StringFlag{
			Name:     "output",
			Aliases:  []string{"o"},
			Value:    "",
			Usage:    "sy|sm|sd|sw|sfull|ly|lm|ld|la|lfull",
			Required: true,
		},
	}

	app.Action = func(c *cli.Context) error {
		day := c.String("day")
		output := c.String("output")
		type1 := "plain"

		inDay := carbon.Time2Carbon(time.Now())
		if day != "" {
			inDay = carbon.Parse(day)
		}

		return Dprint(inDay, output, type1)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// Dprint doit检测工具的输出
func Dprint(inDay carbon.Carbon, output, type1 string) error {
	switch output {
	case "sy":
		fmt.Printf("%s", inDay.ToFormatString("Y"))
	case "sm":
		fmt.Printf("%s", inDay.ToFormatString("m"))
	case "sd":
		fmt.Printf("%s", inDay.ToFormatString("d"))
	case "sw":
		fmt.Printf("%s", inDay.ToFormatString("l"))
	case "ly":
		fmt.Printf("%s", inDay.ToChineseYearString())
	case "lm":
		fmt.Printf("%s", inDay.ToLunarMonthString())
	case "ld":
		fmt.Printf("%s", inDay.ToLunarDayString())
	case "la":
		fmt.Printf("%s", inDay.AnimalYear())
	case "lfull":
		fmt.Printf("%s (%s) 年%s月%s日", inDay.ToChineseYearString(), inDay.AnimalYear(), inDay.ToLunarMonthString(), inDay.ToLunarDayString())
	case "sfull":
		fmt.Printf("%s", inDay.ToDateString())
	default:
		return fmt.Errorf("%s is not one of <sy|sm|sd|sw|ly|lm|ld|la|lfull|sfull>", output)
	}

	return nil

}
