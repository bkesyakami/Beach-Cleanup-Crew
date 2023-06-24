# Beach Cleanup Crew
# File: beach_cleanup.rb

# Create constants
MAX_WATER_LEVEL = 20
MAX_PLASTIC_IN_BAG = 25

# Variables
cleanup_bags = 0
total_trash_collected = 0

# Prompt user for water level
puts "Please enter the current water level in feet."
water_level = gets.chomp.to_i

# Out of bounds check
if water_level > MAX_WATER_LEVEL
  puts "That water level is too high. Please use caution and try again later."
  exit
end

# Loop through beach area
while water_level > 0
  # Collect trash
  puts "Collecting trash in the area."
  piece_of_trash = rand(10)
  puts "Collected #{piece_of_trash} piece(s) of trash."
  total_trash_collected += piece_of_trash
  
  # Add to bag if necessary
  if total_trash_collected >= MAX_PLASTIC_IN_BAG
    puts "Bag is full. Please seal the bag and begin a new one."
    cleanup_bags += 1
    total_trash_collected = 0
  end
  
  # Subtract from water level
  puts "Decreasing water level by one foot."
  water_level -= 1
  puts "Current water level: #{water_level} feet."
  
  # Check for bottom
  if water_level == 0
    puts "Reached the bottom of the beach."
  end
end

# Output
puts "Beach cleanup complete. #{cleanup_bags} bag(s) of trash collected."