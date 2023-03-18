from schema import SiteCoverage
from dotenv import load_dotenv
from schema import SiteCoverageReponse

import uuid
import os
import subprocess


load_dotenv()

RF_SERVER_PATH = os.getenv('RF_SERVER_PATH')
SDF_FILES_PATH = os.getenv('SDF_FILES_PATH')
OUTPUT_PATH = os.getenv('OUTPUT_PATH')

class SitesCoverage:

  @staticmethod
  def calculateCoverage(sites: list[SiteCoverage]):
    print(RF_SERVER_PATH)
    print(SDF_FILES_PATH)
    print(OUTPUT_PATH)
    for site in sites:
        params = ""
        if site.sdf_filename:
            params = params + " -sdf " + SDF_FILES_PATH + site.sdf_filename
        if site.lat:
            params = params + " -lat " + str(site.lat)
        if site.lon:
            params = params + " -lon " + str(site.lon)
        if site.txh:
            params = params + " -txh " + str(site.txh)
        if site.freq:
            params = params + " -f " + str(site.freq)
        if site.erp:
            params = params + " -erp " + str(site.erp)
        if site.rt:
            params = params + " -rt " + str(site.rt)
        if site.dbm:
            params = params + " -dbm "
        params = params + " -m "
        if site.radius:
            params = params + " -R " + str(site.radius)
        if site.res:
            params = params + " -res " + str(site.res)
        if site.pm:
            params = params + " -pm " + str(site.pm)
        
        outputFilePath = ""
        unique_filename = str(uuid.uuid4())
        if site.output_filename:
            outputFilePath = OUTPUT_PATH+site.output_filename
        else:
            outputFilePath = OUTPUT_PATH+unique_filename

        params = params + " -o " + outputFilePath + " 2>&1"
        output = ""
        try:
            subprocess.check_output("mkdir -p " + OUTPUT_PATH, shell=True)

            rfFindCovCommand = RF_SERVER_PATH + "/src/signalserver " + params
            print ("Running "+rfFindCovCommand)
            result = subprocess.check_output(rfFindCovCommand, shell=True, text=True)
            output = result.split("|") 
            print ("Converting to PNG...")
            rfConvertCommand = "convert " + outputFilePath + ".ppm " + outputFilePath + ".png"
            print ("Running "+rfConvertCommand)
            subprocess.check_output(rfConvertCommand, shell=True)
            
            print(output)
        except subprocess.CalledProcessError as e:
            print(e.output)
            return e.output

        if len(output[1:-1]) == 4:
            pngPath = outputFilePath + ".png"
            return SiteCoverageReponse(north_cordinate=output[1], east_cordinate=output[2], south_cordinate=output[3], west_cordinate=output[4], png_path=pngPath)
            
        return output