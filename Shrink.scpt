FasdUAS 1.101.10   ��   ��    k             l     ��  ��    !  Shrink selected HTML file.     � 	 	 6   S h r i n k   s e l e c t e d   H T M L   f i l e .   
  
 l     ��  ��    * $ Created by Stefan Braun 2024-07-08.     �   H   C r e a t e d   b y   S t e f a n   B r a u n   2 0 2 4 - 0 7 - 0 8 .      l     ��  ��    / ) Copyright (c) 2024. All rights reserved.     �   R   C o p y r i g h t   ( c )   2 0 2 4 .   A l l   r i g h t s   r e s e r v e d .      l     ��������  ��  ��        l    � ����  O     �    Q    �     k    �        r     ! " ! l    #���� # 1    ��
�� 
DTsl��  ��   " o      ���� 0 this_selection      $ % $ Z     & '���� & =    ( ) ( o    ���� 0 this_selection   ) J    ����   ' R    �� *��
�� .ascrerr ****      � **** * m     + + � , , 8 P l e a s e   s e l e c t   s o m e   c o n t e n t s .��  ��  ��   %  -�� - X   ! � .�� / . k   1 � 0 0  1 2 1 r   1 8 3 4 3 c   1 6 5 6 5 n   1 4 7 8 7 1   2 4��
�� 
ppth 8 o   1 2���� 0 this_record   6 m   4 5��
�� 
TEXT 4 o      ���� 0 thepath thePath 2  9 : 9 r   9 < ; < ; m   9 : = = � > > , / U s e r s / s b / t m p / s h r u n k e n < o      ���� 0 
targetpath 
targetPath :  ? @ ? l  = =��������  ��  ��   @  A B A l  = =�� C D��   C h b Delete the target folder if it exists already: we want only the newly created file in the folder.    D � E E �   D e l e t e   t h e   t a r g e t   f o l d e r   i f   i t   e x i s t s   a l r e a d y :   w e   w a n t   o n l y   t h e   n e w l y   c r e a t e d   f i l e   i n   t h e   f o l d e r . B  F G F n  = C H I H I   > C�� J���� 0 deletefolder deleteFolder J  K�� K o   > ?���� 0 
targetpath 
targetPath��  ��   I  f   = > G  L M L l  D D��������  ��  ��   M  N O N l  D D�� P Q��   P + % Execute shrinkr for current document    Q � R R J   E x e c u t e   s h r i n k r   f o r   c u r r e n t   d o c u m e n t O  S T S r   D Y U V U b   D U W X W b   D O Y Z Y b   D K [ \ [ m   D E ] ] � ^ ^ p / U s e r s / s b / . l o c a l / b i n / s h r i n k r   s h r i n k   - - n o s t a t s   - - o u t p a t h   \ l  E J _���� _ n   E J ` a ` 1   F J��
�� 
strq a o   E F���� 0 
targetpath 
targetPath��  ��   Z m   K N b b � c c    X l  O T d���� d n   O T e f e 1   P T��
�� 
strq f o   O P���� 0 thepath thePath��  ��   V o      ���� 0 strcmd strCmd T  g h g r   Z e i j i I  Z a�� k��
�� .sysoexecTEXT���     TEXT k o   Z ]���� 0 strcmd strCmd��   j o      ���� 0 	strresult 	strResult h  l m l l  f f��������  ��  ��   m  n o n l  f f�� p q��   p - ' Get the name of the shrunken document	    q � r r N   G e t   t h e   n a m e   o f   t h e   s h r u n k e n   d o c u m e n t 	 o  s t s r   f p u v u n  f l w x w I   g l�� y���� 0 getfilename getFilename y  z�� z o   g h���� 0 
targetpath 
targetPath��  ��   x  f   f g v o      ���� 0 thefile theFile t  { | { r   q ~ } ~ } l  q z ����  b   q z � � � b   q v � � � o   q r���� 0 
targetpath 
targetPath � m   r u � � � � �  / � o   v y���� 0 thefile theFile��  ��   ~ o      ���� 0 newfilepath newFilePath |  � � � l   ��������  ��  ��   �  � � � l   �� � ���   � #  Import the shrunken document    � � � � :   I m p o r t   t h e   s h r u n k e n   d o c u m e n t �  � � � r    � � � � I   ��� ���
�� .DTpacd01DTrc       utxt � o    ����� 0 newfilepath newFilePath��   � o      ���� 0 thenewrecord theNewRecord �  � � � l  � ���������  ��  ��   �  � � � l  � ��� � ���   � 2 , Copy source URL from original to new record    � � � � X   C o p y   s o u r c e   U R L   f r o m   o r i g i n a l   t o   n e w   r e c o r d �  � � � r   � � � � � l  � � ����� � n   � � � � � 1   � ���
�� 
pURL � o   � ����� 0 this_record  ��  ��   � n       � � � 1   � ���
�� 
pURL � o   � ����� 0 thenewrecord theNewRecord �  � � � l  � ���������  ��  ��   �  � � � l  � ��� � ���   � !  Delete the original record    � � � � 6   D e l e t e   t h e   o r i g i n a l   r e c o r d �  � � � I  � ����� �
�� .coredelobool       null��   � �� ���
�� 
DTrc � o   � ����� 0 this_record  ��   �  � � � l  � ���������  ��  ��   �  ��� � I  � ��� � �
�� .DTpacd80bool    ��� utxt � m   � � � � � � �  S h r i n k � �� ���
�� 
info � b   � � � � � m   � � � � � � �  N e w   r e c o r d :   � l  � � ����� � n   � � � � � 1   � ���
�� 
pnam � o   � ����� 0 thenewrecord theNewRecord��  ��  ��  ��  �� 0 this_record   / o   $ %���� 0 this_selection  ��    R      �� � �
�� .ascrerr ****      � **** � o      ���� 0 error_message   � �� ���
�� 
errn � o      ���� 0 error_number  ��    Z  � � � ����� � >  � � � � � l  � � ����� � o   � ����� 0 error_number  ��  ��   � m   � ������� � I  � ��� � �
�� .sysodisAaleR        TEXT � m   � � � � � � �  D E V O N t h i n k � �� � �
�� 
mesS � o   � ����� 0 error_message   � �� ���
�� 
as A � m   � ���
�� EAlTwarN��  ��  ��    5     �� ���
�� 
capp � m     � � � � �  D N t p
�� kfrmID  ��  ��     � � � l     ��������  ��  ��   �  � � � l     �������  ��  �   �  � � � i      � � � I      �~ ��}�~ 0 deletefolder deleteFolder �  ��| � o      �{�{ 0 	thefolder 	theFolder�|  �}   � O     / � � � Q    . � � � � I   �z ��y
�z .coredelobool       null � 4    �x �
�x 
cfol � o   	 
�w�w 0 	thefolder 	theFolder�y   � R      �v � �
�v .ascrerr ****      � **** � o      �u�u 0 error_message   � �t ��s
�t 
errn � o      �r�r 0 error_number  �s   � Z    . � ��q�p � >    � � � l    ��o�n � o    �m�m 0 error_number  �o  �n   � m    �l�l�@ � k    * � �  � � � l   �k � ��k   � 1 + Ignore error if folder does not exist yet.    � � � � V   I g n o r e   e r r o r   i f   f o l d e r   d o e s   n o t   e x i s t   y e t . �  ��j � I   *�i � �
�i .sysodisAaleR        TEXT � m     � � � � � ( D e l e t e   T a r g e t   F o l d e r � �h ��g
�h 
mesS � b    & � � � b    " � � � o     �f�f 0 error_message   � m     ! � � � � �    -   � l  " % ��e�d � c   " % � � � o   " #�c�c 0 error_number   � m   # $�b
�b 
TEXT�e  �d  �g  �j  �q  �p   � m      � ��                                                                                  sevs  alis    \  Macintosh HD               �_�xBD ����System Events.app                                              �����_�x        ����  
 cu             CoreServices  0/:System:Library:CoreServices:System Events.app/  $  S y s t e m   E v e n t s . a p p    M a c i n t o s h   H D  -System/Library/CoreServices/System Events.app   / ��   �  � � � l     �a�`�_�a  �`  �_   �  �  � l     �^�]�\�^  �]  �\    �[ i     I      �Z�Y�Z 0 getfilename getFilename �X o      �W�W 0 thepath thePath�X  �Y   O      r    	 e    

 l   �V�U n     1   
 �T
�T 
pnam n    
 2    
�S
�S 
ditm 4    �R
�R 
cfol o    �Q�Q 0 thepath thePath�V  �U  	 o      �P�P 0 thefilename theFileName m     �                                                                                  sevs  alis    \  Macintosh HD               �_�xBD ����System Events.app                                              �����_�x        ����  
 cu             CoreServices  0/:System:Library:CoreServices:System Events.app/  $  S y s t e m   E v e n t s . a p p    M a c i n t o s h   H D  -System/Library/CoreServices/System Events.app   / ��  �[       �O�O   �N�M�L�N 0 deletefolder deleteFolder�M 0 getfilename getFilename
�L .aevtoappnull  �   � **** �K ��J�I�H�K 0 deletefolder deleteFolder�J �G�G   �F�F 0 	thefolder 	theFolder�I   �E�D�C�E 0 	thefolder 	theFolder�D 0 error_message  �C 0 error_number    ��B�A�@�? ��> ��=�<
�B 
cfol
�A .coredelobool       null�@ 0 error_message   �;�:�9
�; 
errn�: 0 error_number  �9  �?�@
�> 
mesS
�= 
TEXT
�< .sysodisAaleR        TEXT�H 0� , *�/j W X  �� ���%��&%l 
Y hU �8�7�6�5�8 0 getfilename getFilename�7 �4�4   �3�3 0 thepath thePath�6   �2�1�2 0 thepath thePath�1 0 thefilename theFileName �0�/�.
�0 
cfol
�/ 
ditm
�. 
pnam�5 � *�/�-�,EE�U �-�,�+�*
�- .aevtoappnull  �   � **** k     �    �)�)  �,  �+   �(�'�&�( 0 this_record  �' 0 error_message  �& 0 error_number   ,�% ��$�#�" +�!� ���� =�� ]� b����� ������� �� ����
!�	 ������
�% 
capp
�$ kfrmID  
�# 
DTsl�" 0 this_selection  
�! 
kocl
�  
cobj
� .corecnte****       ****
� 
ppth
� 
TEXT� 0 thepath thePath� 0 
targetpath 
targetPath� 0 deletefolder deleteFolder
� 
strq� 0 strcmd strCmd
� .sysoexecTEXT���     TEXT� 0 	strresult 	strResult� 0 getfilename getFilename� 0 thefile theFile� 0 newfilepath newFilePath
� .DTpacd01DTrc       utxt� 0 thenewrecord theNewRecord
� 
pURL
� 
DTrc
� .coredelobool       null
� 
info
� 
pnam
� .DTpacd80bool    ��� utxt�
 0 error_message  ! ���
� 
errn� 0 error_number  �  �	��
� 
mesS
� 
as A
� EAlTwarN� 
� .sysodisAaleR        TEXT�* �)���0 � �*�,E�O�jv  	)j�Y hO ��[��l kh  ��,�&E�O�E�O)�k+ O��a ,%a %�a ,%E` O_ j E` O)�k+ E` O�a %_ %E` O_ j E` O�a ,_ a ,FO*a �l Oa a a  _ a !,%l "[OY�sW &X # $�a % a &a '�a (a )a * +Y hU ascr  ��ޭ