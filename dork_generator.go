package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "sync"
)

func loadDorks(filePath string) ([]string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var dorks []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line != "" {
            dorks = append(dorks, line)
        }
    }
    return dorks, scanner.Err()
}

func generateDorks(domain string, dorksTemplate []string) []string {
    var wg sync.WaitGroup
    dorksChan := make(chan string, len(dorksTemplate))
    
    for _, dork := range dorksTemplate {
        wg.Add(1)
        go func(d string) {
            defer wg.Done()
            customized := strings.ReplaceAll(d, "{domain}", domain)
            dorksChan <- customized
        }(dork)
    }

    go func() {
        wg.Wait()
        close(dorksChan)
    }()

    var generatedDorks []string
    for dork := range dorksChan {
        generatedDorks = append(generatedDorks, dork)
    }
    return generatedDorks
}

func saveToFile(dorks []string, domain string) error {
    filename := fmt.Sprintf("%s_dorks.txt", domain)
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := bufio.NewWriter(file)
    for _, dork := range dorks {
        fmt.Fprintln(writer, dork)
    }
    writer.Flush()
    fmt.Printf("Dorks saved to %s\n", filename)
    return nil
}

func main() {
    // Load dorks
    dorksTemplate, err := loadDorks("dorks.txt")
    if err != nil {
        fmt.Printf("Error: %v. Please create dorks.txt with your dorks.\n", err)
        return
    }
    if len(dorksTemplate) == 0 {
        fmt.Println("No dorks loaded. Exiting.")
        return
    }

    // Get domain
    fmt.Print("Enter the domain (e.g., google.com): ")
    var domain string
    fmt.Scanln(&domain)
    domain = strings.TrimSpace(domain)

    // Generate dorks
    dorks := generateDorks(domain, dorksTemplate)

    // Print dorks
    fmt.Println("\nGenerated Google Dorks:")
    for i, dork := range dorks {
        fmt.Printf("%d. %s\n", i+1, dork)
    }

    // Save option
    fmt.Print("\nDo you want to save these dorks to a file? (yes/no): ")
    var saveOption string
    fmt.Scanln(&saveOption)
    saveOption = strings.ToLower(saveOption)

    if saveOption == "yes" || saveOption == "y" {
        if err := saveToFile(dorks, domain); err != nil {
            fmt.Printf("Error saving file: %v\n", err)
        }
    }
}
