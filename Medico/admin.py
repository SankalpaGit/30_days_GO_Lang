from django.contrib import admin
from .models import Product
from .models import Message, RequestBlood, DonateBlood


# Register your models here.
admin.site.register(Product)
admin.site.register(Message)
admin.site.register(RequestBlood)
admin.site.register(DonateBlood)