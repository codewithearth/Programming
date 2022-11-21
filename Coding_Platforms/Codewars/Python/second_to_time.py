def format_duration(seconds): # convert seconds to accurate time
    
    if seconds == 0:# if second is zero then time is "now"
        return "now"

    else:#convert time into minutes,hours,days and years
        minute, second = seconds/60, seconds%60
        hour, minute = minute//60, minute%60
        day, hour = hour//24, hour%24
        year, day = day//365, day%365

    #create empty list to store all data
    queue_data = []

    queue_data.append(queue(int(year), "year"))
    queue_data.append(queue(int(day),"day"))
    queue_data.append(queue(int(hour),"hour"))
    queue_data.append(queue(int(minute),"minute"))
    queue_data.append(queue(int(second),"second"))

    #remove all empty spaces from list
    queue_data = [ data for data in queue_data if data!= ""]

    if len(queue_data)>1:# if data is more then one then add " and " at last two vlues
        queue_data[-2] = queue_data[-2]+" and "+ queue_data[-1]
        queue_data.remove(queue_data[-1])
        
    #join all the result by ","
    result = ", ".join(queue_data)
    return result

def queue(count,name): # count data values

    if count == 1: # if data value is one then add one and data name
        return "1 "+name

    elif count > 1:# if data value is more then one then add count of data and add "s" in the end of data
        return str(count)+" "+name+"s"
    return ""
    
print(format_duration(162))