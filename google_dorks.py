# Google Dorks Automation Tool
# Reads dorks from an external file and replaces {domain} with user-provided domain
# Uses multithreading for dork generation
# Save results in a file named after the domain (e.g., google.com_dorks.txt)

import concurrent.futures

def load_dorks(file_path="dorks.txt"):
    """Load Google Dorks from an external text file."""
    try:
        with open(file_path, "r") as file:
            # Read lines, strip whitespace, and filter out empty lines
            dorks = [line.strip() for line in file if line.strip()]
        return dorks
    except FileNotFoundError:
        print(f"Error: {file_path} not found. Please create it with your dorks.")
        return []

def replace_domain(dork, domain):
    """Helper function to replace {domain} in a single dork."""
    return dork.replace("{domain}", domain)

def generate_dorks(domain, dorks_template):
    """Generate Google Dorks using multithreading."""
    generated_dorks = []
    
    # Use ThreadPoolExecutor for parallel processing
    with concurrent.futures.ThreadPoolExecutor() as executor:
        # Map the replace_domain function across all dorks
        futures = [executor.submit(replace_domain, dork, domain) for dork in dorks_template]
        # Collect results as they complete
        generated_dorks = [future.result() for future in concurrent.futures.as_completed(futures)]
    
    # Maintain original order (optional, since order might not matter)
    return generated_dorks

def save_to_file(dorks, domain):
    """Save the generated dorks to a text file named after the domain."""
    filename = f"{domain}_dorks.txt"
    with open(filename, "w") as file:
        for dork in dorks:
            file.write(dork + "\n")
    print(f"Dorks saved to {filename}")

def main():
    # Load dorks from external file
    dorks_template = load_dorks("dorks.txt")
    
    if not dorks_template:
        print("No dorks loaded. Exiting.")
        return
    
    # Get domain from user
    domain = input("Enter the domain (e.g., google.com): ").strip()
    
    # Generate dorks with multithreading
    dorks = generate_dorks(domain, dorks_template)
    
    # Print dorks
    print("\nGenerated Google Dorks:")
    for i, dork in enumerate(dorks, 1):
        print(f"{i}. {dork}")
    
    # Ask if user wants to save to file
    save_option = input("\nDo you want to save these dorks to a file? (yes/no): ").lower()
    if save_option == "yes" or save_option == "y":
        save_to_file(dorks, domain)

if __name__ == "__main__":
    main()
