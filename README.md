# grpc_calculator

# we have 5 api endpoints in which we will 5 different operations by applying gRPC framework 


# Sum Api & Sqare_Root Api (This Is Called Uniary Api's)

in these 2  api's endpoint we will just send a single request to server and server will send a single response back to client
   
like ---sending one request to server --->>> Then <<< Getting A Single Response From Server

                      Step : 01
           --->->--- Request 01 -->-->--- 
Client---|                                |---Server
           ---<-<--- Response 02 --<--<---    
                      Step : 02
                      
                      
                      
                   
# Primes ( This Type Of Api's Called Server Streaming Api's , In Which Server Will Send A Stream And Client Will Recieve It )


   in these api endpoint we will just send a single request to server and server will send multiple response back to client
   
   like ---sending one request to server --->>> Then <<< Getting Multiple Responses From Server
   
   
                   Step : 01
             --->->---Request-->-->--- 
Client---|                                |---Server
           ---<-<---Response 01 --<--<--- 
           ---<-<---Response 02 --<--<--- 
           ---<-<---Response 03 --<--<--- 
           ---<-<---Response 04 --<--<--- 
                     Step : 02
                   
                   
# Average (This Type Of Api's Called Client Straming Api's , In Which Client Will Send A Stream And Server Will Recieve It )
in these type of api endpoints we will send multiple requests to server and server will a send a single response back to client
   
sending multiple request to server --->>> Then <<< Getting Single Response From Server



                       Step : 01
                       
             --->->---Request 01 -->-->---
             --->->---Request 02 -->-->---
             --->->---Request 03 -->-->---
             --->->---Request 04 -->-->---
Client---|                                   |---Server
             ---<-<---Response 01 --<--<--- 
           
                       Step : 02
                     
                     
                     
# Maximum ( This Type Of Api's Called Bi_Directional Api's In Which Both Client And Server Will Send And Recieve Stream's Of Each Other )
(This Type Of Api's Called Client Straming Api's)
 
in these type of api endpoints we will just send multiple requests to server and server will send multiple response back to client simenteneously (req/resp)
   
sending multiple request to server --->>> Then <<< Getting Multiple Responses From Server simenteneously
                      
                      Step : 01
                      
             --->->---Request 01 -->-->---
             --->->---Request 02 -->-->---
             --->->---Request 03 -->-->---
             --->->---Request 04 -->-->---
Client---|          simenteneously            |---Server
             ---<-<---Response 01 --<--<--- 
             ---<-<---Response 02 --<--<--- 
             ---<-<---Response 03 --<--<--- 
             ---<-<---Response 04 --<--<--- 
                      
                      
                      Step : 02
                   
